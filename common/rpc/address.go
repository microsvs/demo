package rpc

import (
	"fmt"

	"github.com/microsvs/base"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
)

var (
	PROVINCE_NAME__QUERY__SCHEMA = `
	query {
		province(id: %d){
			name
		}
	}
	`

	CITY_NAME__QUERY__SCHEMA = `
	query {
		city(id: %d){
			name
		}
	}
	`
)

func GetProvinceNameByIdRPC(id int) (name string, err error) {
	var respMap map[string]interface{}
	query := fmt.Sprintf(PROVINCE_NAME__QUERY__SCHEMA, id)
	if respMap, err = rpc.CallService(nil, base.Service2Url(rpc.FGSAddress), query); err != nil {
		log.ErrorRaw("[GetProvinceNameByIdRPC] failed to get address , err=%s", err.Error())
		return "", errors.FGEHTTPRPCError
	}
	temp := respMap["province"].(interface{})
	tempMap := temp.(map[string]interface{})
	return tempMap["name"].(string), nil
}

func GetCityNameByIdRPC(id int) (name string, err error) {
	var respMap map[string]interface{}
	query := fmt.Sprintf(CITY_NAME__QUERY__SCHEMA, id)
	if respMap, err = rpc.CallService(nil, base.Service2Url(rpc.FGSAddress), query); err != nil {
		log.ErrorRaw("[GetCityNameByIdRPC] failed to get address , err=%s", err.Error())
		return "", errors.FGEHTTPRPCError
	}
	temp := respMap["city"].(interface{})
	tempMap := temp.(map[string]interface{})
	return tempMap["name"].(string), nil
}
