package rpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/microsvs/base"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"github.com/microsvs/base/pkg/utils"
	common "github.com/microsvs/demo/common/types"
)

var (
	// 用户基本信息查询
	UserQueryByIdSchema = `
	query{
		user(userid: "?"){
			id
			nickname
			phone
			user_status
		}
	}
	`
	UserQueryByPhoneSchema = `
	query{
		user(phone: "?"){
			id
			nickname
			phone
			user_status
		}
	}
	`
)

func GetUserByUserIdRPC(ctx context.Context, userId string) (
	user *common.User, err error) {
	if len(userId) <= 0 {
		return nil, fmt.Errorf("param `user_id` empty")
	}
	var (
		data map[string]interface{}
	)
	user = new(common.User)
	query := strings.Replace(UserQueryByIdSchema, "?", userId, 1)
	data, err = rpc.CallService(ctx, base.Service2Url(rpc.FGSUser), query)
	if err != nil {
		log.ErrorRaw("[GetUserByUserIdRPC] http rpc failed. err=%s", err.Error())
		return nil, errors.FGEHTTPRPCError
	}
	base.FixTypeFromGoToGraphql(data["user"], common.GLUser)
	if err = utils.Decode(data, "user", user); err != nil {
		log.ErrorRaw("[GetUserByUserIdRPC] decode user data failed. err=%s", err.Error())
		return nil, errors.FGEDataParseError
	}
	return
}

func GetUserByMobileRPC(ctx context.Context, phone string) (
	user *common.User, err error) {
	if len(phone) <= 0 {
		return nil, fmt.Errorf("param `phone` empty")
	}
	var data map[string]interface{}
	user = new(common.User)
	query := strings.Replace(UserQueryByPhoneSchema, "?", phone, 1)
	if data, err = rpc.CallService(ctx, base.Service2Url(rpc.FGSUser), query); err != nil {
		log.ErrorRaw("[GetUserByUserIdRPC] http rpc failed. err=%s", err.Error())
		return nil, errors.FGEHTTPRPCError
	}
	base.FixTypeFromGoToGraphql(data["user"], common.GLUser)
	if err = utils.Decode(data, "user", user); err != nil {
		log.ErrorRaw("[GetUserByUserIdRPC] decode user data failed. err=%s", err.Error())
		return nil, errors.FGEDataParseError
	}
	return
}
