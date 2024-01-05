package response

type Response struct {
	Statuscode int64
	Message    string
	Data       interface{}
	Error      interface{}
}

func Responses(code int64, message string, data interface{}, error interface{}) Response {
	return Response{
		Statuscode: code,
		Message:    message,
		Data:       data,
		Error:      error,
	}
}
