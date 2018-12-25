package main

import (
	"context"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base"
	"github.com/microsvs/base/cmd/cache"
	pcache "github.com/microsvs/base/pkg/cache"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"github.com/microsvs/base/pkg/types"
	"github.com/microsvs/base/pkg/utils"
	itypes "github.com/microsvs/demo/common/types"
)

var (
	USER_SCHEMA string = `
		 query{
			 %s(mobile: "%s"){
				 id
				 nickname
				 mobile
				 city_ids
				 user_type
				 advertisers{
					 advertiser_id
					 name
				 }
				 status
			 }
		 }
		`
	USER_ID_SCHEMA string = `
		 query{
			 %s(user_id: "%s"){
				 id
				 nickname
				 mobile
				 city_ids
				 user_type
				 advertisers{
					 advertiser_id
					 name
				 }
				 status
			 }
		 }
		`
	TOKEN__SCHEMA string = `
			mutation{
				new_token(user_id: "%s"){
					token
					token_expire
					refresh_token
					refresh_token_expire
				}
			}
		`
)

func newToken(ctx context.Context, user *itypes.User) (token *types.Token, err error) {
	var data map[string]interface{}
	if data, err = rpc.CallService(
		ctx,
		base.Service2Url(rpc.FGSToken),
		fmt.Sprintf(TOKEN__SCHEMA, user.ID),
	); err != nil {
		log.ErrorRaw("[queryToken]: failed to request newToken by user[%v], err is %v.", user.ID, err)
		return
	}
	token = new(types.Token)
	if err = utils.Decode(data, "new_token", token); err != nil {
		log.ErrorRaw("[queryToken] decode token failed. err=%s", err.Error())
		return
	}
	return
}

func queryToken(p graphql.ResolveParams) (token *types.Token, err error) {
	var (
		data          map[string]interface{}
		TOKEN__SCHEMA = `
		query{
			query_token{
				token
				token_expire
				refresh_token
				refresh_token_expire
				user_id
			}
		}
		`
	)

	if data, err = rpc.CallService(p.Context, base.Service2Url(rpc.FGSToken), TOKEN__SCHEMA); err != nil {
		log.ErrorRaw("[queryToken] rpc http request failed. err=%s", err)
		return
	}
	fmt.Printf("token: %v\n", data)
	token = new(types.Token)
	if err = utils.Decode(data, "query_token", token); err != nil {
		log.ErrorRaw("[queryToken] decode data failed. err=%s", err)
		return
	}
	return
}

func me(p graphql.ResolveParams) (interface{}, error) {
	var (
		me    Me
		err   error
		ok    bool
		buser *types.User
		data  map[string]interface{}
	)
	if buser, ok = p.Context.Value(rpc.KeyUser).(*types.User); !ok {
		log.ErrorRaw("[gateway:me]: failed to get user.")
		return nil, errors.FGEInvalidUserID
	}
	me.Basic = new(itypes.User)
	if data, err = rpc.CallService(nil, base.Service2Url(rpc.FGSUser),
		fmt.Sprintf(USER_ID_SCHEMA, "me", buser.ID)); err != nil {
		return nil, err
	}
	data["me"] = base.FixTypeFromGraphqlToGo(data["me"], itypes.GLUser)
	if err = utils.Decode(data, "me", me.Basic); err != nil {
		log.ErrorRaw("[me] decode data %v failed, err=%s", data["me"], err.Error())
		return nil, err
	}
	if me.Token, err = queryToken(p); err != nil {
		log.ErrorRaw("[me] get token failed. err=%s", err.Error)
		return nil, err
	}
	return &me, nil
}

func verifyCode(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSUser)
}

func matchVerifyCode(mobile string, code string) (err error) {
	var (
		conn      pcache.Connection
		matchCode string
	)
	if conn = cache.SlaveCache(rpc.FGSUser); conn == nil {
		log.ErrorRaw("[login]: failed to connect redis.")
		return errors.FGECacheError
	}
	defer conn.Close()
	if matchCode, err = redis.String(conn.Get("demo/user/sms/" + mobile)); err != nil {
		if err != redis.ErrNil {
			log.ErrorRaw("[login]: failed to get verify code, err is %v.", err)
			return errors.FGECacheError
		}
		return errors.FGEExpiredVerifyCode
	}
	if code != matchCode {
		log.ErrorRaw("[login]: request code=%s, redis code=%s.", code, matchCode)
		return errors.FGEInvalidVerifyCode
	}
	return
}

func login(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok           bool
		err          error
		code, mobile string
		user         Me
		data         map[string]interface{}
	)
	if code, ok = p.Args["code"].(string); !ok {
		log.ErrorRaw("[login] param `code` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	if mobile, ok = p.Args["mobile"].(string); !ok {
		log.ErrorRaw("[login] param `mobile` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	method := "query_user_by_mobile"
	data, err = rpc.CallService(nil, base.Service2Url(rpc.FGSUser), fmt.Sprintf(USER_SCHEMA, method, mobile))
	if err != nil {
		log.ErrorRaw("[login] rpc call service failed. err=%s", err.Error())
		return nil, err
	}
	if err = matchVerifyCode(mobile, code); err != nil {
		return nil, err
	}
	user.Basic = new(itypes.User)
	base.FixTypeFromGraphqlToGo(data[method], itypes.GLUser)
	if err = utils.Decode(data, method, user.Basic); err != nil {
		log.ErrorRaw("[login] decode user failed. err=%s", err.Error())
		return nil, errors.FGEDataParseError
	}
	// 强退其他设备登陆态
	rpc.CallService(p.Context, base.Service2Url(rpc.FGSToken), `mutation{ logout }`)

	if user.Token, err = newToken(p.Context, user.Basic); err != nil {
		log.ErrorRaw("[gateway:login]: failed to request token by user[%v], err is %v.", user.Basic.ID, err)
		return nil, err
	}
	return &user, nil
}
