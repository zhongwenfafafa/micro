package response

type ErrorResponse struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code int32 `json:"code"`
	Message string	`json:"message"`
	Data interface{} `json:"data"`
}

func NewSuccessRes(
	code int32, msg string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Code: code,
		Message: msg,
		Data: data,
	}
}

func NewErrorRes(
	code int32, msg string) *ErrorResponse {
	return &ErrorResponse{
		Code: code,
		Message: msg,
	}
}