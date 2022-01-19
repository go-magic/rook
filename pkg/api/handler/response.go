package handler

type Response struct {
	Error string
	Data  interface{}
}

func NewResponse(err string, data interface{}) *Response {
	return &Response{
		Error: err,
		Data:  data,
	}
}
