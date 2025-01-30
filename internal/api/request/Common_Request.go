package request

// 用户登录时，前端传输参数需要绑定的相关结构体
type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
