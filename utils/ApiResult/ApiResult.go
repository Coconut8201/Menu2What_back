package ApiResult

type ApiResult struct {
	Result  bool   `json:"result"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
