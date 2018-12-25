package rpc

import (
	"fmt"
	"reflect"

	"github.com/microsvs/base"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"github.com/microsvs/demo/common/utils"
)

var (
	UPLOAD__PICTURE__SCHEMA = `
	mutation{
		add_pictures(urls: %s)
	}
	`

	QUERY__PICTURE__SCHEMA = `
	query{
		query_picture(id: %d)
	}
	`

	QUERY__PICTURES__SCHEMA = `
	query{
		query_pictures(ids: %v)
	}
	`
)

func UploadPicturesRPC(urls []string) (ids []int, err error) {
	var respMap map[string]interface{}
	query := fmt.Sprintf(UPLOAD__PICTURE__SCHEMA, utils.ConvertToGraphqlList(urls))
	if respMap, err = rpc.CallService(nil, base.Service2Url(rpc.FGSImage), query); err != nil {
		log.ErrorRaw("[UploadPicturesRPC]: failed to upload picture, err= %s.", err.Error())
		return nil, errors.FGEHTTPRPCError
	}
	tmps := respMap["add_pictures"].([]interface{})
	for _, tmp := range tmps {
		ids = append(ids, int(reflect.ValueOf(tmp).Float()))
	}
	fmt.Printf("ids=%v\n", ids)
	return
}

func GetPictureRPC(id int) (url string, err error) {
	var respMap map[string]interface{}
	query := fmt.Sprintf(QUERY__PICTURE__SCHEMA, id)
	if respMap, err = rpc.CallService(nil, base.Service2Url(rpc.FGSImage), query); err != nil {
		log.ErrorRaw("[GetPictureRPC] failed to get picture , err=%s", err.Error())
		return "", errors.FGEHTTPRPCError
	}
	temp := respMap["query_picture"].(interface{})
	return reflect.ValueOf(temp).String(), nil
}

func GetPicturesRPC(ids []int) (urls []string, err error) {
	var respMap map[string]interface{}
	query := fmt.Sprintf(QUERY__PICTURES__SCHEMA, ids)
	if respMap, err = rpc.CallService(nil, base.Service2Url(rpc.FGSImage), query); err != nil {
		log.ErrorRaw("[GetPicturesRPC] failed to get pictures, err=%s", err)
		return nil, errors.FGEHTTPRPCError
	}
	tmps := respMap["query_pictures"].([]interface{})
	for _, tmp := range tmps {
		urls = append(urls, reflect.ValueOf(tmp).String())
	}
	return
}
