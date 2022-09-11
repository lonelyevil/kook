package kook

import (
	"encoding/json"
	"strconv"
)

// RestError is the error type for errors from kook
type RestError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// Error provides the formatted error string
func (r RestError) Error() string {
	return "[" + strconv.Itoa(r.Code) + "] " + r.Message
}

func newRestErrorFromGeneralResp(r *EndpointGeneralResponse) error {
	return &RestError{Code: r.Code, Message: r.Message, Data: r.Data}
}
