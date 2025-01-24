package ApiResult

type ApiResult struct {
	Result  bool        `json:"result"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResult(code int, data interface{}) *ApiResult {
	return &ApiResult{
		Result: true,
		Code:   code,
		Data:   data,
	}
}

func NewFailResult(code int, message string) *ApiResult {
	return &ApiResult{
		Result:  false,
		Code:    code,
		Message: message,
	}
}
