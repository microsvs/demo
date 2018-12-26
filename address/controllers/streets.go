package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/demo/address/models"
)

func Streets(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["district_id"].(int); !ok {
		log.ErrorRaw("[Streets] param `district_id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetStreetsByDistrictId(id, nil)
}
func Street(p graphql.ResolveParams) (interface{}, error) {
	var (
		ok bool
		id int
	)
	if id, ok = p.Args["street_id"].(int); !ok {
		log.ErrorRaw("[Street] param `street_id` empty")
		return nil, errors.FGEInvalidRequestParam
	}
	return models.GetStreetById(id, nil)
}
