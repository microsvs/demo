package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base"
	"github.com/microsvs/base/pkg/rpc"
)

func logout(p graphql.ResolveParams) (interface{}, error) {
	return base.RedirectRequest(p, rpc.FGSToken)
}
