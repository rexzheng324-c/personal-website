package result

const (
	Success = "200"
	Fail    = "500"

	// business status code
	ErrorUsernameUsed   = "1001"
	ErrorPasswordWrong  = "1002"
	ErrorUserNotExist   = "1003"
	ErrorTokenExist     = "1004"
	ErrorTokenRuntime   = "1005"
	ErrorTokenWrong     = "1006"
	ErrorTokenTypeWrong = "1007"
	ErrorUserNoRight    = "1008"
)

var codeMsg = map[string]string{
	Success: "Success",

	// business status code
	ErrorUsernameUsed:   "用户名已存在！",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在,请重新登陆",
	ErrorTokenRuntime:   "TOKEN已过期,请重新登陆",
	ErrorTokenWrong:     "TOKEN不正确,请重新登陆",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登陆",
	ErrorUserNoRight:    "该用户无权限",
}

func GetErrMsg(code string) string {
	return codeMsg[code]
}
