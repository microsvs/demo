package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/demo/token/models"
	"github.com/microsvs/demo/token/types"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name:        "Query",
				Description: "查询所有Token相关的信息",
				Fields: graphql.Fields{
					"token": &graphql.Field{
						Type: types.GLToken,
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type:        graphql.String,
								Description: "通过token获得所有token关联信息",
							},
						},
						Resolve: models.QueryToken,
					},
					"query_token": &graphql.Field{
						Type:        types.GLToken,
						Description: "查询token",
						Resolve:     models.TokenByUserId,
					},
				},
			}),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name:        "Mutation",
				Description: "更新token操作",
				Fields: graphql.Fields{
					"new_token": &graphql.Field{
						Type:        types.GLToken,
						Description: "获取新token",
						Args: graphql.FieldConfigArgument{
							"user_id": &graphql.ArgumentConfig{
								Type:        graphql.NewNonNull(graphql.String),
								Description: "用户ID",
							},
						},
						Resolve: models.NewToken,
					},
					"logout": &graphql.Field{
						Type:        graphql.Boolean,
						Description: "退出登录",
						Resolve:     models.Logout,
					},
				},
			}),
	})
