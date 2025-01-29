package common

// 各种信息类的结构体，将所有的返回信息类的接口模块化返回的不同类型都使用data一个接口进行覆写处理
type Result struct {
	Code int         `json:"code"` //这个结果是否是成功的
	Data interface{} `json:"data"` //结果应该包含的数据，是一个抽象接口可以进行多次覆写
	Msg  string      `json:"msg"`  //相关的信息
}

type PageResult struct {
	Total   int64       `json:"total"`   //总记录页数
	Records interface{} `json:"records"` //当前页数数据集合，主要用于搜索
}
