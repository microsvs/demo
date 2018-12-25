package types

import "github.com/graphql-go/graphql"

var FGError = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Error",
		Description: "返回错误码",
		Fields: graphql.Fields{
			"err_code": &graphql.Field{
				Type:        graphql.Int,
				Description: "错误码",
			},
			"err_msg": &graphql.Field{
				Type:        graphql.String,
				Description: "错误信息",
			},
		},
	},
)
