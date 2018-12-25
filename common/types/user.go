package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/microsvs/demo/common/consts"
)

//User 用户信息
type User struct {
	ID        string           `db:"user_id" json:"id"`
	Mobile    string           `db:"mobile" json:"mobile"`
	NickName  string           `db:"nickname" json:"nickname" db:"nickname"`
	UserType  consts.USER_TYPE `db:"user_type" json:"user_type"`
	CityIds   string           `db:"city_ids" json:"city_ids"`
	Status    consts.STATUS    `db:"status" json:"status"`
	UpdatedAt time.Time        `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time        `db:"created_at" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}

type Advertiser struct {
	AdvertiserId string        `db:"advertiser_id" json:"advertiser_id"`
	Name         string        `db:"name" json:"name"`
	Status       consts.STATUS `db:"status" json:"status"`
	UpdatedAt    time.Time     `db:"updated_at" json:"updated_at"`
	CreatedAt    time.Time     `db:"created_at" json:"created_at"`
	PictureUrls  []string      `db:"-" json:"picture_urls"`
}

func (Advertiser) TableName() string {
	return "advertisers"
}

//GLUser GraphQL 用户类型
var GLUser = graphql.NewObject(GLUserConfig)

//GLUserConfig GraphQL 用户配置
var GLUserConfig = graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "用户ID",
		},
		"nickname": &graphql.Field{
			Type:        graphql.String,
			Description: "用户昵称",
		},
		"mobile": &graphql.Field{
			Type:        graphql.String,
			Description: "电话号码，数据库保存加密以后的电话号码，返回脱敏以后的电话号码",
		},
		"city_ids": &graphql.Field{
			Type:        graphql.String,
			Description: "城市列表，以';'分割",
		},
		"user_type": &graphql.Field{
			Type:        GLUserType,
			Description: "用户类型：10：销售；20：广告客户",
		},
		"status": &graphql.Field{
			Type:        GLUserStatus,
			Description: "用户状态",
		},
	},
}

/*
GLUserStatus ... 用户使用状态
*/
var GLUserStatus = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserStatus",
	Description: "用户状态",
	Values: graphql.EnumValueConfigMap{
		"Normal": &graphql.EnumValueConfig{
			Value:       consts.STATUS__OK,
			Description: "正常状态",
		},
		"Delete": &graphql.EnumValueConfig{
			Value:       consts.STATUS__INVALID,
			Description: "被删除状态",
		},
	},
})

var GLUserType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserType",
	Description: "用户类型",
	Values: graphql.EnumValueConfigMap{
		"Saler": &graphql.EnumValueConfig{
			Value:       consts.USER_TYPE__SALER,
			Description: "销售员",
		},
		"Advertiser": &graphql.EnumValueConfig{
			Value:       consts.USER_TYPE__ADVERTISER,
			Description: "广告客户",
		},
	},
})

var GLSex = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "Sex",
		Description: "性别",
		Values: graphql.EnumValueConfigMap{
			"Male": &graphql.EnumValueConfig{
				Value:       consts.MALE,
				Description: "男",
			},
			"Female": &graphql.EnumValueConfig{
				Value:       consts.FEMALE,
				Description: "女",
			},
			"Unknown": &graphql.EnumValueConfig{
				Value:       consts.UNKNOWN,
				Description: "未知性别",
			},
		},
	},
)
