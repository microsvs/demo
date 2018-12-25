package types

import "github.com/graphql-go/graphql"

type OS int16

const (
	Ios     OS = 10
	Android OS = 20
)

var GLOS = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "OS",
		Description: "操作系统类型: IOS | Android",
		Values: graphql.EnumValueConfigMap{
			"ios": &graphql.EnumValueConfig{
				Value:       Ios,
				Description: "ios",
			},
			"android": &graphql.EnumValueConfig{
				Value:       Android,
				Description: "android",
			},
		},
	},
)
