package serializer

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error error       `json:"error,omitempty"`
}

func DataResponse(data interface{}) Response {
	return Response{Data: data}
}
