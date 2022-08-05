package kore

import (
	"github.com/gcp-iot/model"
)

type successResponse struct {
	Message string      `json:"message"  validate:"required"`
	Data    interface{} `json:"data"  validate:"required"`
}
type failResponse struct {
	Message string `json:"message"  validate:"required"`
	Error   string `json:"error"  validate:"required"`
}

func FrameSuccess(statusCode int, msg string, data interface{}) model.Response {

	frame := successResponse{Message: msg, Data: data}
	return model.Response{StatusCode: statusCode, Message: frame}
}
func FrameError(statusCode int, msg string, err string) model.Response {
	frame := failResponse{Message: msg, Error: err}
	return model.Response{StatusCode: statusCode, Message: frame}
}
