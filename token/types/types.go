package types

import (
	"github.com/graphql-go/graphql"
)

//GLTokenConfig GraphQL Token配置
var GLTokenConfig = graphql.ObjectConfig{
	Name: "Token",
	Fields: graphql.Fields{
		"user_id": &graphql.Field{
			Type:        graphql.String,
			Description: "Token对应的用户ID",
		},
		"token": &graphql.Field{
			Type:        graphql.String,
			Description: "Token",
		},
		"token_expire": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Token过期时间",
		},
		"refresh_token": &graphql.Field{
			Type:        graphql.String,
			Description: "使用RefreshToken进行刷新操作",
		},
		"refresh_token_expire": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "RefreshToken的过期时间，如果RefreshToken已过期则需要重新登录",
		},
	},
}

//GLToken GraphQL Token 定义
var GLToken = graphql.NewObject(GLTokenConfig)
