package result

type box struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessBox(data interface{}) *box {
	return &box{
		Code:    Success,
		Message: GetErrMsg(Success),
		Data:    data,
	}
}
