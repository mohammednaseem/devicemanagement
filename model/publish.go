package model

type PublishCreate struct {
	Operation string       `json:"operation" validate:"required"`
	Data      DeviceCreate `json:"data" validate:"required"`
}
type PublishUpdate struct {
	Operation string       `json:"operation" validate:"required"`
	Data      DeviceUpdate `json:"data" validate:"required"`
}
type PublishDelete struct {
	Operation string       `json:"operation" validate:"required"`
	Data      DeviceDelete `json:"data" validate:"required"`
}
