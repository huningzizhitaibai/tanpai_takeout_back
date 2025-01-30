package model

// 简化表，只记录基本数据
type User_basic struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     int    `json:"type"`      //用户的权限0-管理员，1-用户，2-商户，3-骑手
	RealName string `json:"real_name"` //真实姓名，非必须
}
