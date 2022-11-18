package common

type UnifiedResponse struct {
	Code    string
	Message string
	Request string
	Data    interface{}
}

func (u *UnifiedResponse) InitCodeAndMessageForSuccess() {
	u.Code = "00000"
	u.Message = "success"
}

func NewUnifiedResponse(code, message, request string) *UnifiedResponse {
	return &UnifiedResponse{
		Code:    code,
		Message: message,
		Request: request,
	}
}

func NewUnifiedResponseEmpty() *UnifiedResponse {
	var u = UnifiedResponse{}
	u.InitCodeAndMessageForSuccess()
	return &u
}

func NewUnifiedResponseMessgae(message string) *UnifiedResponse {
	var u = UnifiedResponse{}
	u.InitCodeAndMessageForSuccess()
	u.Message = message
	return &u
}

func NewUnifiedResponseMessgaeData(message string, data interface{}) *UnifiedResponse {
	var u = UnifiedResponse{}
	u.InitCodeAndMessageForSuccess()
	u.Data = data
	u.Message = message
	return &u
}

func NewUnifiedResponseData(data interface{}) *UnifiedResponse {
	var u = UnifiedResponse{Data: data}
	u.InitCodeAndMessageForSuccess()
	return &u
}
