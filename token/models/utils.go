package models

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
	pcache "github.com/microsvs/base/cmd/cache"
	icache "github.com/microsvs/base/pkg/cache"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"github.com/microsvs/base/pkg/timer"
	"github.com/microsvs/base/pkg/types"
)

var (
	prefixTokenPath = "demo/token/%s"
	prefixUserPath  = "demo/user/%s"
)

func GetTokenInfoByToken(token string) (result *types.Token, err error) {
	var (
		conn icache.Connection
		bts  []byte
	)
	if len(token) <= 0 {
		log.ErrorRaw("[GetUserIdByToken] param `token` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	if conn = pcache.MasterCache(rpc.FGSToken); conn == nil {
		log.ErrorRaw("[GetUserIdByToken] get redis connection failed.")
		return nil, errors.FGECacheError
	}
	if bts, err = redis.Bytes(conn.Get(fmt.Sprintf(prefixTokenPath, token))); err != nil {
		if err != redis.ErrNil {
			log.ErrorRaw("[GetUserIdByToken] get token failed from redis, err=%s", err.Error())
			return nil, errors.FGECacheError
		}
		return nil, nil
	}
	result = new(types.Token)
	if len(bts) > 0 {
		if err = json.Unmarshal(bts, result); err != nil {
			log.ErrorRaw("[GetUserIdByToken] json unmarshal failed. err=%s", err.Error())
			return nil, errors.FGEDataParseError
		}
	}
	return
}

func SaveCacheToken(token *types.Token) (err error) {
	fmt.Println("enter SaveCacheToken")
	defer fmt.Println("left SaveCacheToken")
	var (
		conn   icache.Connection
		bts    []byte
		expire int
	)
	if len(token.Token) <= 0 || len(token.UserId) <= 0 {
		log.ErrorRaw("[SaveCacheToken] param `token||userid` empty")
		return errors.FGEInvalidRequestParam
	}
	if conn = pcache.MasterCache(rpc.FGSToken); conn == nil {
		log.ErrorRaw("[SaveCacheToken] get redis connection failed.")
		return errors.FGECacheError
	}
	expire = int(token.TokenExpire.Sub(timer.Now).Seconds())
	if bts, err = json.Marshal(*token); err != nil {
		log.ErrorRaw("[SaveCacheToken] json marshal failed. err=%s", err.Error())
		return errors.FGEDataParseError
	}
	if err = conn.Set(fmt.Sprintf(prefixTokenPath, token.Token), bts); err != nil {
		log.ErrorRaw("[SaveCacheToken] set redis key failed. err=%s", err.Error())
		return errors.FGECacheError
	}
	if err = conn.Expire(fmt.Sprintf(prefixTokenPath, token.Token), expire); err != nil {
		log.ErrorRaw("[SaveCacheToken] expire redis key failed. err=%s", err.Error())
		return errors.FGECacheError
	}
	return
}

func DelCacheToken(token string) (err error) {
	var (
		conn icache.Connection
	)
	if len(token) <= 0 {
		log.ErrorRaw("[DelCacheToken] param `token` empty")
		return errors.FGEInvalidRequestParam
	}
	if conn = pcache.MasterCache(rpc.FGSToken); conn == nil {
		log.ErrorRaw("[DelCacheToken] get redis connection failed.")
		return errors.FGECacheError
	}
	if err = conn.Del(fmt.Sprintf(prefixTokenPath, token)); err != nil {
		log.ErrorRaw("[DelCacheToken] delete redis key failed. err=%s", err.Error())
		return errors.FGECacheError
	}
	return
}

func DelCacheUser(userId string) (err error) {
	var (
		conn icache.Connection
	)
	if len(userId) <= 0 {
		log.ErrorRaw("[DelCacheUser] param `user_id` empty")
		return errors.FGEInvalidRequestParam
	}
	if conn = pcache.MasterCache(rpc.FGSToken); conn == nil {
		log.ErrorRaw("[DelCacheUser] get redis connection failed.")
		return errors.FGECacheError
	}
	if err = conn.Del(fmt.Sprintf(prefixUserPath, userId)); err != nil {
		log.ErrorRaw("[DelCacheUser] delete redis key failed. err=%s", err.Error())
		return errors.FGECacheError
	}
	return
}
