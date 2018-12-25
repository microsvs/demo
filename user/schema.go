package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base/pkg/types"
	itypes "github.com/microsvs/demo/common/types"
	"github.com/microsvs/demo/user/controllers"
)

var schema graphql.Schema

func init() {
	var err error
	if schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query,
			Mutation: mutation,
		},
	); err != nil {
		panic(err)
	}
}

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Query",
		Description: "【灵动】接口查询列表",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        types.GLUser,
				Description: "基础用户信息，内部使用",
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "用户ID",
					},
				},
				Resolve: controllers.QueryBasicUserById,
			},
			"me": &graphql.Field{
				Type:        itypes.GLUser,
				Description: "查询 我 相关的所有信息",
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "用户ID",
					},
				},
				Resolve: controllers.QueryUserById,
			},
			"query_user_by_mobile": &graphql.Field{
				Type:        itypes.GLUser,
				Description: "用户登录",
				Args: graphql.FieldConfigArgument{
					"mobile": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "手机号码，目前仅支持国内手机",
					},
				},
				Resolve: controllers.QueryUserByMobile,
			},
		},
	},
)

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Description: "【灵动】接口修改列表",
	Fields: graphql.Fields{
		"verify_code": &graphql.Field{
			Type:        graphql.Boolean, // the return type for this field
			Description: "请求短信验证码",
			Args: graphql.FieldConfigArgument{
				"mobile": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "手机号码，目前仅支持国内手机可以不含国家代码，也可以+86开头",
				},
			},
			Resolve: controllers.VerifyCode,
		},
	},
})
