package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/demo/address/models"
)

func GetCities(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["province_id"].(int); !ok {
		log.ErrorRaw("[GetCities] param `province_id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetCitiesByProvinceId(id, nil)
}

func GetCity(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["id"].(int); !ok {
		log.ErrorRaw("[GetCity] param `id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetCityById(id, nil)
}
