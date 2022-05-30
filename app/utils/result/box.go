package result

type successBox struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessBox(data interface{}) *successBox {
	return &successBox{
		Code:    Success,
		Message: GetErrMsg(Success),
		Data:    data,
	}
}

type failBox struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewFailBox(code string, err error) *failBox {
	return &failBox{
		Code:    code,
		Message: err.Error(),
	}
}
