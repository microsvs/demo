package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base"
	"github.com/microsvs/base/pkg/rpc"
)

func getProvinces(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func getProvince(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func getCities(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func getCity(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func districts(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func district(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func streets(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}

func street(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSAddress)
}
