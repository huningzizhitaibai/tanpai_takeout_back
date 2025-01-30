package enum

//不同用户类型

type UserType int

const (
	controller UserType = 0
	user       UserType = 1
	shop       UserType = 2
	deliver    UserType = 3
	error      UserType = -1
)

const (
	User       string = "User"
	Controller string = "Controller"
	Shop       string = "Shop"
	Deliver    string = "Deliver"
	Error      string = "Error"
)

func StrT2IntT(st string) UserType {
	switch st {
	case User:
		return user
	case Controller:
		return controller
	case Shop:
		return shop
	case Deliver:
		return deliver
	default:
		return error
	}
}

func IntT2StrT(it UserType) string {
	switch it {
	case user:
		return User
	case controller:
		return Controller
	case shop:
		return Shop
	case deliver:
		return Deliver
	default:
		return Error
	}
}
