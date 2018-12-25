package types

import "github.com/graphql-go/graphql"

var GLProvince = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Province",
		Description: "省份信息",
		Fields: graphql.Fields{
			"province_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "省份ID",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "省份名称",
			},
			"citycode": &graphql.Field{
				Type:        graphql.String,
				Description: "直辖市编号",
			},
			"adcode": &graphql.Field{
				Type:        graphql.String,
				Description: "区域编码",
			},
			"longtitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "经度",
			},
			"latitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "纬度",
			},
			"weight_value": &graphql.Field{
				Type:        graphql.Int,
				Description: "权重",
			},
		},
	},
)

var GLCity = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "City",
		Description: "城市信息",
		Fields: graphql.Fields{
			"city_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "城市ID",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "城市名称",
			},
			"citycode": &graphql.Field{
				Type:        graphql.String,
				Description: "城市编码",
			},
			"province_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "省份ID",
			},
			"longtitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "经度",
			},
			"latitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "纬度",
			},
			"weight_value": &graphql.Field{
				Type:        graphql.Int,
				Description: "权重",
			},
		},
	},
)
var GLDistrict = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "District",
		Description: "区/县信息",
		Fields: graphql.Fields{
			"district_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "区/县ID",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "区/县名称",
			},
			"citycode": &graphql.Field{
				Type:        graphql.String,
				Description: "城市编码",
			},
			"adcode": &graphql.Field{
				Type:        graphql.String,
				Description: "区域编码",
			},
			"province_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "省份ID",
			},
			"city_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "城市ID",
			},
			"longtitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "经度",
			},
			"latitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "纬度",
			},
			"weight_value": &graphql.Field{
				Type:        graphql.Int,
				Description: "权重",
			},
		},
	},
)
var GLStreet = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Street",
		Description: "街道信息",
		Fields: graphql.Fields{
			"street_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "街道ID",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "街道名称",
			},
			"citycode": &graphql.Field{
				Type:        graphql.String,
				Description: "城市编码",
			},
			"adcode": &graphql.Field{
				Type:        graphql.String,
				Description: "区域编码",
			},
			"province_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "省份ID",
			},
			"city_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "城市ID",
			},
			"district_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "区/县ID",
			},
			"longtitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "经度",
			},
			"latitude": &graphql.Field{
				Type:        graphql.Float,
				Description: "纬度",
			},
			"weight_value": &graphql.Field{
				Type:        graphql.Int,
				Description: "权重",
			},
		},
	},
)
