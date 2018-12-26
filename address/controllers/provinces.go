package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/demo/address/models"
)

func GetProvinces(p graphql.ResolveParams) (interface{}, error) {
	return models.GetProvinces(nil)
}

func GetProvince(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["id"].(int); !ok {
		log.ErrorRaw("[GetProvince] param `id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetProvinceById(id, nil)
}
