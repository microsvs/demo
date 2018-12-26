package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/demo/address/controllers"
	common "github.com/microsvs/demo/common/types"
)

var schema graphql.Schema

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Query",
		Description: "查询地址(省份、城市、县/区和街道)信息",
		Fields: graphql.Fields{
			"provinces": &graphql.Field{
				Type:        graphql.NewList(common.GLProvince),
				Description: "省份列表",
				Resolve:     controllers.GetProvinces,
			},
			"province": &graphql.Field{
				Type:        common.GLProvince,
				Description: "具体省份信息",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "省份ID",
					},
				},
				Resolve: controllers.GetProvince,
			},
			"cities": &graphql.Field{
				Type:        graphql.NewList(common.GLCity),
				Description: "城市列表",
				Args: graphql.FieldConfigArgument{
					"province_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "省份ID",
					},
				},
				Resolve: controllers.GetCities,
			},
			"city": &graphql.Field{
				Type:        common.GLCity,
				Description: "城市信息",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "城市ID",
					},
				},
				Resolve: controllers.GetCity,
			},
			"districts": &graphql.Field{
				Type:        graphql.NewList(common.GLDistrict),
				Description: "区/县列表",
				Args: graphql.FieldConfigArgument{
					"city_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "城市ID",
					},
				},
				Resolve: controllers.Districts,
			},
			"district": &graphql.Field{
				Type:        common.GLDistrict,
				Description: "区/县",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "区/县ID",
					},
				},
				Resolve: controllers.District,
			},
			"streets": &graphql.Field{
				Type:        graphql.NewList(common.GLStreet),
				Description: "街道列表",
				Args: graphql.FieldConfigArgument{
					"district_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "区/县ID",
					},
				},
				Resolve: controllers.Streets,
			},
			"street": &graphql.Field{
				Type:        common.GLStreet,
				Description: "街道信息",
				Args: graphql.FieldConfigArgument{
					"street_id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "街道ID",
					},
				},
				Resolve: controllers.Street,
			},
		},
	},
)

func init() {
	var err error
	schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		},
	)
	if err != nil {
		panic(err)
	}
}
