package controllers

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/graphql-go/graphql"
	pcache "github.com/microsvs/base/cmd/cache"
	"github.com/microsvs/base/cmd/discovery"
	icache "github.com/microsvs/base/pkg/cache"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	itypes "github.com/microsvs/base/pkg/types"
	"github.com/microsvs/demo/common/types"
	"github.com/microsvs/demo/common/utils"
	"github.com/microsvs/demo/user/models"
)

/*
1. 根据用户ID，获取用户基本信息
2. 根据mobile， 获取用户基本信息
3. 根据用户ID，查询用户的所有权限
*/
var (
	//手机号正则
	mobileRxg = regexp.MustCompile(`^1([3578][0-9]|14[57]|5[^4])\d{8}$`)
)

func QueryUserById(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok             bool
		err            error
		userId         string
		user           = new(types.User)
		conn           icache.Connection
		userBts        []byte
		isCacheInvalid bool
	)
	if userId, ok = p.Args["user_id"].(string); !ok {
		log.ErrorRaw("[QueryUser]  param `user_id` empty")
		return nil, errors.FGEInvalidUserID
	}
	if conn = pcache.MasterCache(rpc.FGSUser); conn == nil {
		log.ErrorRaw("[QueryUser] get redis connection failed.")
		return nil, errors.FGECacheError
	}
	if userBts, err = redis.Bytes(conn.Get("demo/user/" + userId)); err != nil {
		log.ErrorRaw("[QueryUser] get cache data failed. err=%s", err.Error())
		if redis.ErrNil != err {
			return nil, errors.FGECacheError
		}
		isCacheInvalid = true
	}
	// cache hit
	if !isCacheInvalid {
		if err = json.Unmarshal(userBts, user); err != nil {
			log.ErrorRaw("[QueryUser] json unmarshal user failed in cache, err=%s", err.Error())
			return nil, errors.FGEDataParseError
		}
		return user, nil
	}
	if user, err = models.GetUserByUserId(userId, nil); err != nil {
		log.ErrorRaw("[QueryUser] get user info failed. err=%s", err.Error())
		return nil, err
	}
	var bts []byte
	if bts, err = json.Marshal(*user); err != nil {
		log.ErrorRaw("QueryUser] json marshal failed. err=%s", err.Error())
		return nil, errors.FGEDataParseError
	}
	if err = conn.Set("demo/user/"+userId, bts); err != nil {
		log.ErrorRaw("[QueryUser] set redis key failed. err=%s", err.Error())
		return nil, errors.FGECacheError
	}
	if err = conn.Expire("demo/user/"+userId, 30*24*60*60); err != nil {
		// 一周有效期
		log.ErrorRaw("[QueryUser] expire redis key failed. err=%s", err.Error())
		return nil, errors.FGECacheError
	}
	return user, nil
}

// 可不同类型的用户使用: B端、C端和OP运营端
func VerifyCode(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok     bool
		mobile string
		err    error
		conn   icache.Connection
	)
	// expire verify code in redis
	expiredTs, _ := strconv.Atoi(discovery.KVRead("tool/user/sms_expired_ts", "300"))
	if expiredTs <= 60 {
		log.ErrorRaw("[VerifyCode] read sms verify code in redis expire is too small")
		return false, errors.FGEZKConfigError
	}
	// param checks
	if mobile, ok = p.Args["mobile"].(string); !ok {
		log.ErrorRaw("[VerifyCode] param `mobile` is empty")
		return false, errors.FGEInvalidRequestParam
	}
	if _, err = models.GetUserByMobile(mobile, nil); err != nil {
		return false, err
	}
	if ok = mobileRxg.MatchString(mobile); !ok {
		log.ErrorRaw("[VerifyCode] regex match mobile failed.")
		return false, errors.FGEInvalidMobile
	}
	vCode := utils.GenerateVerifyCode()
	if conn = pcache.MasterCache(rpc.FGSUser); conn == nil {
		log.ErrorRaw("[VerifyCode] get redis connection failed.")
		return false, errors.FGECacheError
	}
	if err = conn.Set("demo/user/sms/"+mobile, vCode); err != nil {
		log.ErrorRaw("[VerifyCode] set key in redis failed, err=%s", err.Error())
		return false, errors.FGECacheError
	}
	if err = conn.Expire("demo/user/sms/"+mobile, expiredTs); err != nil {
		log.ErrorRaw("[VerifyCode] expire key in redis failed. err=%s", err.Error())
		return false, errors.FGECacheError
	}
	if _, err = models.SendShortMessage(mobile, vCode); err != nil {
		log.ErrorRaw("[VerifyCode] send verify code failed.", err.Error())
		return false, err
	}

	return true, nil
}

func QueryUserByMobile(p graphql.ResolveParams) (interface{}, error) {
	var (
		mobile string
		ok     bool
		user   *types.User
		err    error
	)
	if mobile, ok = p.Args["mobile"].(string); !ok {
		log.ErrorRaw("[QueryUserBymobile] param `mobile` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	if user, err = models.GetUserByMobile(mobile, nil); err != nil {
		return nil, err
	}
	return user, nil
}

func QueryBasicUserById(p graphql.ResolveParams) (interface{}, error) {
	var (
		userId string
		user   *itypes.User
		ok     bool
		err    error
	)
	if userId, ok = p.Args["user_id"].(string); !ok {
		log.ErrorRaw("[QueryUser]  param `user_id` empty")
		return nil, errors.FGEInvalidUserID
	}
	if user, err = models.GetBasicUserById(userId, nil); err != nil {
		log.ErrorRaw("[QueryUser] get user info failed. err=%s", err.Error())
		return nil, err
	}
	return user, nil
}
