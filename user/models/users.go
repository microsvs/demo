package models

import (
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	itypes "github.com/microsvs/base/pkg/types"
	"github.com/microsvs/demo/common/types"
	udb "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

func GetUserByUserId(userId string, db sqlbuilder.Database) (user *types.User, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSUser); db == nil {
			log.ErrorRaw("[GetUserByUserId] get mysql instance failed. ")
			return nil, errors.FGEDBError
		}
	}
	user = new(types.User)
	sess := db.Collection(user.TableName())
	if err = sess.Find("user_id", userId).One(user); err != nil {
		if err == udb.ErrNoMoreRows {
			log.ErrorRaw("[GetUserByUserId] user `%s` not exist in base_user_infos", userId)
			return nil, errors.FGEInvalidUserID
		}
		log.ErrorRaw("[GetUserByUserId] read `base_user_infos` record failed. err=%s", err.Error())
		return nil, err
	}
	return
}

func GetBasicUserById(userId string, db sqlbuilder.Database) (user *itypes.User, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSUser); db == nil {
			log.ErrorRaw("[GetUserByUserId] get mysql instance failed. ")
			return nil, errors.FGEDBError
		}
	}
	user = new(itypes.User)
	sess := db.Collection(user.TableName())
	if err = sess.Find("user_id", userId).One(user); err != nil {
		if err == udb.ErrNoMoreRows {
			log.ErrorRaw("[GetUserByUserId] user `%s` not exist in base_user_infos", userId)
			return nil, errors.FGEInvalidUserID
		}
		log.ErrorRaw("[GetUserByUserId] read `base_user_infos` record failed. err=%s", err.Error())
		return nil, err
	}
	return
}

func GetUserByMobile(mobile string, db sqlbuilder.Database) (user *types.User, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSUser); db == nil {
			log.ErrorRaw("[QueryUserBymobile] get mysql instance failed.")
			return nil, errors.FGEDBError
		}
	}
	user = new(types.User)
	sess := db.Collection(user.TableName())
	if err = sess.Find("mobile", mobile).One(user); err != nil {
		if err == udb.ErrNoMoreRows {
			return nil, errors.FGEInvalidMobile
		}
		log.ErrorRaw("[QueryUserBymobile]  read `user` failed. err: %s", err.Error())
		return nil, errors.FGEDBError
	}
	return
}
