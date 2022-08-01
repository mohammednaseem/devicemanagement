package model

type PublishDeviceCreate struct {
	Operation string       `json:"operation" validate:"required"`
	Entity    string       `json:"entity" validate:"required"`
	Data      DeviceCreate `json:"data" validate:"required"`
}
type PublishDeviceUpdate struct {
	Operation string       `json:"operation" validate:"required"`
	Entity    string       `json:"entity" validate:"required"`
	Data      DeviceUpdate `json:"data" validate:"required"`
}
type PublishDeviceDelete struct {
	Operation string       `json:"operation" validate:"required"`
	Entity    string       `json:"entity" validate:"required"`
	Data      DeviceDelete `json:"data" validate:"required"`
}
type PublishRegistryCreate struct {
	Operation string         `json:"operation" validate:"required"`
	Entity    string         `json:"entity" validate:"required"`
	Data      RegistryCreate `json:"data" validate:"required"`
}
type PublishRegistryUpdate struct {
	Operation string         `json:"operation" validate:"required"`
	Entity    string         `json:"entity" validate:"required"`
	Data      RegistryUpdate `json:"data" validate:"required"`
}
type PublishRegistryDelete struct {
	Operation string         `json:"operation" validate:"required"`
	Entity    string         `json:"entity" validate:"required"`
	Data      RegistryDelete `json:"data" validate:"required"`
}
