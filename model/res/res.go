package res

type Response struct {
	// 状态码
	Code int `json:"code"`
	// 消息
	Msg string `json:"msg"`
	// 数据
	Data interface{} `json:"data"`
	// 错误信息
	Error string `json:"error"`
}
