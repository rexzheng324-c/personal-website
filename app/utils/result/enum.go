package result

const (
	Success = "200"
	Fail    = "500"

	ParamsError    = "400"
	RecordNotFound = "404"

	// user
	UserNameExist    = "user-10001"
	UserNameNotExist = "user-10002"
	WrongPassword    = "user-10003"
	NotAdmin         = "user-10004"
)

var codeMsg = map[string]string{
	Success: "Success",
	Fail:    "Fail",
}

func GetErrMsg(code string) string {
	return codeMsg[code]
}
