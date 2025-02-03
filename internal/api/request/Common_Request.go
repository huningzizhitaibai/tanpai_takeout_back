package request

// 用户登录时，前端传输参数需要绑定的相关结构体
type LoginDTO struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

type SignUpDTO_User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpDTO_Shop struct {
	Username           string `json:"username" form:"username" binding:"required"`
	Password           string `json:"password" form:"password" binding:"required"`
	RealName           string `json:"realname" form:"realname" binding:"required"`
	IDNumber           string `json:"idnumber" form:"idnumber" binding:"required"`
	CertificateForFood string `json:"certificateForFood" form:"certificateForFood" binding:"required"` //食品安全证
	IDCard1            string `json:"id_card1" form:"id_card1" binding:"required"`                     //法人身份证正反两面
	IDCard2            string `json:"id_card2" form:"id_card2" binding:"required"`
	CertificateForShop string `json:"certificateForShop" form:"certificateForShop" binding:"required"` //店铺许可
}

type SignupDTO_Deliver struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	RealName    string `json:"realname" form:"realname" binding:"required"`
	IDNumber    string `json:"idnumber" form:"idnumber" binding:"required"`
	IDCard1     string `json:"id_card1" form:"id_card" binding:"required"`
	IDCard2     string `json:"id_card2" form:"id_card2" binding:"required"`
	StudentCard string `json:"student_card" form:"student_card"`
	IsStudent   bool   `json:"is_student" form:"is_student"`
}

type SignupDTO_Controller struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	RealName   string `json:"realname" form:"realname" binding:"required"`
	IDNumber   string `json:"idnumber" form:"idnumber" binding:"required"`
	IDCard1    string `json:"id_card1" form:"id_card1" binding:"required"`
	IDCard2    string `json:"id_card2" form:"id_card2" binding:"required"`
	InviteCode string `json:"invite_code" form:"invite_code" binding:"required"` //平台的管理员必须填写邀请码进行注册，确保管理员的安全性
}
