package models

import (
	"time"

	"github.com/graphql-go/graphql"
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"github.com/microsvs/base/pkg/timer"
	"github.com/microsvs/base/pkg/types"
	uuid "github.com/satori/go.uuid"
	udb "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

func QueryToken(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok        bool
		token     string
		err       error
		db        sqlbuilder.Database
		tokenInfo *types.Token
	)
	if token, ok = p.Args["token"].(string); !ok {
		log.ErrorRaw("[QueryToken] param `token` empty")
		return nil, errors.FGEInvalidToken
	}
	if tokenInfo, err = GetTokenInfoByToken(token); err != nil {
		log.ErrorRaw("[QueryToken] get token from redis failed. err=%s", err.Error())
		return nil, errors.FGECacheError
	}
	// non-nil return
	if tokenInfo != nil && len(tokenInfo.Token) > 0 {
		return tokenInfo, nil
	}
	if db = pdb.SlaveDB(rpc.FGSToken); db == nil {
		log.ErrorRaw("[queryToken] get mysql connection failed.")
		return nil, errors.FGEDBError
	}
	tokenInfo = new(types.Token)
	sess := db.Collection(tokenInfo.TableName())
	if err = sess.Find().And(
		udb.Cond{
			"token":          token,
			"token_expire >": timer.Now.Format("2006-01-02"),
		},
	).One(tokenInfo); err != nil {
		if err == udb.ErrNoMoreRows {
			return nil, errors.FGEInvalidToken
		}
		log.ErrorRaw("[QueryToken] get mysql record failed. err=%s", err.Error())
		return nil, errors.FGEDBError
	}
	if err = SaveCacheToken(tokenInfo); err != nil {
		log.ErrorRaw("[QueryToken]  save redis token failed. err=%s", err.Error())
		return nil, errors.FGECacheError
	}
	return tokenInfo, nil
}

func NewToken(p graphql.ResolveParams) (interface{}, error) {
	var (
		err     error
		userId  string
		ok      bool
		token   types.Token
		db      sqlbuilder.Database
		isExist bool = true // if exist token, delete it
	)
	if userId, ok = p.Args["user_id"].(string); !ok {
		log.ErrorRaw("[NewToken] param `user_id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	if db = pdb.MasterDB(rpc.FGSToken); db == nil {
		log.ErrorRaw("[NewToken] get mysql connection failed.")
		return nil, errors.FGEDBError
	}
	sess := db.Collection(token.TableName())
	if err = sess.Find("user_id", userId).One(&token); err != nil {
		if err != udb.ErrNoMoreRows {
			log.ErrorRaw("[NewToken] get mysql record failed. err=%s", err.Error())
			return nil, errors.FGEDBError
		}
		isExist = false
	}
	if isExist {
		if err = DelCacheToken(token.Token); err != nil {
			log.ErrorRaw("[NewToken] delete cache token failed, err=%s", err.Error())
			return nil, errors.FGECacheError
		}
		if err = sess.Find("user_id", userId).Delete(); err != nil {
			log.ErrorRaw("[NewToken] delete token record failed, err=%s", err.Error())
			return nil, errors.FGEDBError
		}
	}
	v4Token := uuid.NewV4()
	v4RefreshToken := uuid.NewV4()
	token = types.Token{
		UserId:             userId,
		Token:              v4Token.String(),
		TokenExpire:        timer.Now.Add(24 * 30 * time.Hour),
		RefreshToken:       v4RefreshToken.String(),
		RefreshTokenExpire: timer.Now.Add(24 * 150 * time.Hour),
	}
	if _, err = sess.Insert(&token); err != nil {
		log.ErrorRaw("[NewToken] create mysql record failed. err=%s", err.Error())
		return nil, errors.FGEDBError
	}
	if err = SaveCacheToken(&token); err != nil {
		log.ErrorRaw("[NewToken] save cache token failed. err=%s", err.Error())
		return nil, errors.FGECacheError
	}
	return &token, nil
}

func Logout(p graphql.ResolveParams) (interface{}, error) {
	var (
		user  *types.User
		err   error
		ok    bool
		token = new(types.Token)
		db    sqlbuilder.Database
	)
	if user, ok = p.Context.Value(rpc.KeyUser).(*types.User); !ok {
		log.ErrorRaw("[Logout] get user from context.")
		return false, errors.FGEInvalidToken
	}
	// del mysql
	if db = pdb.MasterDB(rpc.FGSToken); db == nil {
		log.ErrorRaw("[Logout] init mysql connection failed.")
		return false, errors.FGEDBError
	}
	sess := db.Collection(token.TableName())
	res := sess.Find("user_id", user.ID)
	if err = res.One(token); err != nil {
		if err == udb.ErrNoMoreRows {
			return true, nil
		}
		log.ErrorRaw("[Logout] read `tokens` record failed. err=%s", err)
		return false, err
	}
	if err = res.Delete(); err != nil {
		log.ErrorRaw("[Logout] delete `tokens` record failed. err=%s", err)
		return false, err
	}
	// del cache
	if err = DelCacheToken(token.Token); err != nil {
		log.ErrorRaw("[Logout] delete token cache failed. err=%s", err)
		return false, err
	}
	if err = DelCacheUser(user.ID); err != nil {
		log.ErrorRaw("[DelCacheToken] delete cache user failed. err=%s", err.Error())
		return false, err
	}
	return true, nil
}

func TokenByUserId(p graphql.ResolveParams) (interface{}, error) {
	var (
		user  *types.User
		ok    bool
		err   error
		token *types.Token
		db    sqlbuilder.Database
	)
	if user, ok = p.Context.Value(rpc.KeyUser).(*types.User); !ok {
		log.ErrorRaw("[TokenByUserId] get user from context failed.")
		return nil, errors.FGEInvalidToken
	}
	if db = pdb.SlaveDB(rpc.FGSToken); db == nil {
		log.ErrorRaw("[TokenByUserId] init mysql connection failed.")
		return nil, errors.FGEDBError
	}
	token = new(types.Token)
	sess := db.Collection(token.TableName())
	if err = sess.Find().And(
		udb.Cond{
			"user_id":         user.ID,
			"token_expire > ": timer.Now.Format("2006-01-02"),
		},
	).One(token); err != nil {
		log.ErrorRaw("[QueryToken] get mysql record failed. err=%s", err.Error())
		return nil, errors.FGEDBError
	}
	return token, nil
}
