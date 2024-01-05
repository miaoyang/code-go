package common

const (
	HTTP_OK     = 200
	SERVER_FAIL = 500

	EMPTY          = 20000
	FAIL           = 20001
	VALILD_FAIL    = 20002
	INSERT_DB_FAIL = 20003
	REDIS_SET_FAIL = 20004

	AUTHORIZATION_FAIL        = 20100
	AUTHORIZATION_EMPTY       = 20101
	AUTHORIZATION_TYPE_WRONG  = 20102
	AUTHORIZATION_PARSE_ERROR = 20103
	AUTHORIZATION_EXPIRETIME  = 20104

	USER_NOT_EXIST            = 21000
	USER_PASSWORD_NOT_MATCHED = 21001
)

var GlobalServerMap = map[int]string{
	EMPTY:                     "输入的信息为空",
	FAIL:                      "服务器内部错误",
	VALILD_FAIL:               "校验信息失败",
	INSERT_DB_FAIL:            "插入数据库失败",
	AUTHORIZATION_FAIL:        "用户认证失败",
	USER_NOT_EXIST:            "用户不存在",
	USER_PASSWORD_NOT_MATCHED: "用户密码输入错误，不匹配",
}

func GetMapInfo(code int) string {
	s := GlobalServerMap[code]
	return s
}

// R 请求的返回基本信息
type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewR(code int, msg string, data interface{}) *R {
	return &R{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Ok() *R {
	return &R{
		Code: 200,
		Msg:  "success",
		Data: "",
	}
}

func OkWithMsgData(msg string, data interface{}) *R {
	return NewR(HTTP_OK, msg, data)
}

func OkWithData(data interface{}) *R {
	return NewR(HTTP_OK, "success", data)
}

func Fail() *R {
	return &R{
		Code: SERVER_FAIL,
		Msg:  GlobalServerMap[FAIL],
		Data: nil,
	}
}

func FailWithCodeMsgData(code int, msg string, data interface{}) *R {
	return NewR(code, msg, data)
}

func FailWithCodeMsg(code int, msg string) *R {
	return NewR(code, msg, nil)
}

func FailWithMsgData(msg string, data interface{}) *R {
	return NewR(SERVER_FAIL, msg, data)
}

func FailWithMsg(msg string) *R {
	return NewR(SERVER_FAIL, msg, nil)
}

type PageRes struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"count"`
}

func NewPageRes(data interface{}, total int64) *PageRes {
	return &PageRes{
		Data:  data,
		Total: total,
	}
}
