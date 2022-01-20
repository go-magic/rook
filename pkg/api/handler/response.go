package handler

type Response struct {
	Message string
	Data    interface{}
}

func NewResponse(msg string, data interface{}) *Response {
	return &Response{
		Message: msg,
		Data:    data,
	}
}
