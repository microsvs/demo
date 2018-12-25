package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base"
	ptypes "github.com/microsvs/base/pkg/types"
	itypes "github.com/microsvs/demo/common/types"
)

//Me 所有个人信息
type Me struct {
	Basic *itypes.User  `json:"basic"`
	Token *ptypes.Token `json:"token"`
}

//GLMe GraphQL Me Struct
var GLMe = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Me",
		Fields: graphql.Fields{
			"basic": &graphql.Field{
				Type:        itypes.GLUser,
				Description: "查询用户电话,昵称等基本信息",
			},
			"token": &graphql.Field{
				Type:        base.HideGLFields(ptypes.GLToken, "user_id"),
				Description: "查询用户当前Token，RefreshToken等",
			},
		},
	},
)

var schema graphql.Schema

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "query",
		Description: "【灵动】查询接口列表",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type:        GLMe,
				Description: "查询 我 相关的所有信息",
				Resolve:     me,
			},
			"provinces": &graphql.Field{
				Type:        graphql.NewList(itypes.GLProvince),
				Description: "省份列表",
				Resolve:     getProvinces,
			},
			"province": &graphql.Field{
				Type:        itypes.GLProvince,
				Description: "具体省份信息",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "省份ID",
					},
				},
				Resolve: getProvince,
			},
			"cities": &graphql.Field{
				Type:        graphql.NewList(itypes.GLCity),
				Description: "城市列表",
				Args: graphql.FieldConfigArgument{
					"province_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "省份ID",
					},
				},
				Resolve: getCities,
			},
			"city": &graphql.Field{
				Type:        itypes.GLCity,
				Description: "城市信息",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "城市ID",
					},
				},
				Resolve: getCity,
			},
			"districts": &graphql.Field{
				Type:        graphql.NewList(itypes.GLDistrict),
				Description: "区/县列表",
				Args: graphql.FieldConfigArgument{
					"city_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "城市ID",
					},
				},
				Resolve: districts,
			},
			"district": &graphql.Field{
				Type:        itypes.GLDistrict,
				Description: "区/县",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "区/县ID",
					},
				},
				Resolve: district,
			},
			"streets": &graphql.Field{
				Type:        graphql.NewList(itypes.GLStreet),
				Description: "街道列表",
				Args: graphql.FieldConfigArgument{
					"district_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "区/县ID",
					},
				},
				Resolve: streets,
			},
			"street": &graphql.Field{
				Type:        itypes.GLStreet,
				Description: "街道信息",
				Args: graphql.FieldConfigArgument{
					"street_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "街道ID",
					},
				},
				Resolve: street,
			},
			"errors": &graphql.Field{
				Type:        graphql.NewList(itypes.FGError),
				Description: "返回错误码列表",
				Resolve:     retErrors,
			},
		},
	},
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mutation",
		Description: "【DEMO】修改接口列表",
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
				Resolve: verifyCode,
			},
			"login": &graphql.Field{
				Type:        GLMe, // the return type for this field
				Description: "注册或者登录统一接口",
				Args: graphql.FieldConfigArgument{
					"mobile": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "手机号码，目前仅支持国内手机可以不含国家代码，也可以+86开头",
					},
					"code": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "验证码",
					},
				},
				Resolve: login,
			},
			"logout": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "退出登录",
				Resolve:     logout,
			},
		},
	},
)

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
