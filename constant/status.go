package constant

type ErrorStatus struct {
	StatusCode int32
	StatusMsg  string
}

var (
	SUCCESS              = &ErrorStatus{StatusCode: 0, StatusMsg: "Success"}
	ERR_WRONG_PASSWORD   = &ErrorStatus{StatusCode: 2, StatusMsg: "Wrong Password"}
	ERR_RECORD_NOT_FOUND = &ErrorStatus{StatusCode: 4, StatusMsg: "Record Not Found"}
	ERR_SERVICE_INTERNAL = &ErrorStatus{StatusCode: 6, StatusMsg: "Internal Error"}
)

var (
	HTTP_SUCCESS                  = &ErrorStatus{StatusCode: 0, StatusMsg: ""}
	HTTP_ERR_WRONG_PASSWORD       = &ErrorStatus{StatusCode: 2, StatusMsg: "密码错误"}
	HTTP_ERR_RECORD_NOT_FOUND     = &ErrorStatus{StatusCode: 4, StatusMsg: "找不到这条文本"}
	HTTP_ERR_SERVICE_INTERNAL     = &ErrorStatus{StatusCode: 6, StatusMsg: "服务器打瞌睡了"}
	HTTP_ERR_FREQUENCY_OVER_LIMIT = &ErrorStatus{StatusCode: 8, StatusMsg: "操作太快了，请稍后再试"}
	HTTP_ERR_WRONG_PARAMS         = &ErrorStatus{StatusCode: 10, StatusMsg: "参数有误"}
)
