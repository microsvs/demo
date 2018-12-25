package consts

type STATUS int16

const (
	// 记录状态:  10：有效；20：无效
	STATUS__OK      STATUS = 10
	STATUS__INVALID STATUS = 20
)

type USER_TYPE int16

const (
	// 用户类型: 10: 销售；20: 广告主客户
	USER_TYPE__SALER      USER_TYPE = 10
	USER_TYPE__ADVERTISER USER_TYPE = 20
)

// 性别
type SEX_TYPE int16

const (
	MALE    SEX_TYPE = 10 // 男
	FEMALE  SEX_TYPE = 20 // 女
	UNKNOWN SEX_TYPE = 30 // 未知
)
