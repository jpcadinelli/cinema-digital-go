package middleware

import "fmt"

type ResponseBridge struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
}

func NewResponseBridge(err error, data any) *ResponseBridge {
	var errStr string
	if err != nil {
		errStr = fmt.Sprintf("%s", err)
	}
	return &ResponseBridge{
		Error: errStr,
		Data:  data,
	}
}
