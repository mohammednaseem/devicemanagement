package model

type Frame struct {
	StateCode int         `json:"stateCode"  validate:"required"`
	Message   string      `json:"message"  validate:"required"`
	Details   interface{} `json:"details"  validate:"required"`
}

func FrameResponse(statusCode int, msg string, details interface{}) Response {

	frame := Frame{StateCode: 0, Message: msg, Details: details}
	return Response{StatusCode: statusCode, Message: frame}
}
