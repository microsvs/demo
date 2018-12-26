package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/demo/address/models"
)

func Districts(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["city_id"].(int); !ok {
		log.ErrorRaw("[Districts] param `city_id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetDistrictsByCityId(id, nil)
}

func District(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["id"].(int); !ok {
		log.ErrorRaw("[District] param `id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetDistrictById(id, nil)
}
