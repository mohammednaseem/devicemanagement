package model

import (
	cloudiot "google.golang.org/api/cloudiot/v1"
)

/////// request
type RequestRegistry struct {
	UpdateMask               string                              `json:"updatemask" validate:""`
	Parent                   string                              `json:"parent" validate:"required"`
	Id                       string                              `json:"id" validate:"required"`
	Name                     string                              `json:"name" validate:""`
	EventNotificationConfigs []*cloudiot.EventNotificationConfig `json:"eventNotificationConfigs" validate:"required"`
	StateNotificationConfig  *cloudiot.StateNotificationConfig   `json:"stateNotificationConfig"  validate:""`
	MqttConfig               cloudiot.MqttConfig                 `json:"mqttConfig"  validate:""`
	HttpConfig               cloudiot.HttpConfig                 `json:"httpConfig"  validate:""`
	LogLevel                 string                              `json:"logLevel"  validate:""`
	Credentials              []*cloudiot.RegistryCredential      `json:"credentials"  validate:""`
}
type RequestDevice struct {
	UpdateMask  string                       `json:"updatemask" validate:""`
	Parent      string                       `json:"parent" validate:"required"`
	Id          string                       `json:"id" validate:"required"`
	Name        string                       `json:"name" validate:""`
	Credentials []*cloudiot.DeviceCredential `json:"credentials" validate:"required"`
	LogLevel    string                       `json:"logLevel"  validate:""`
	Blocked     bool                         `json:"blocked"  validate:""`
	Metadata    map[string]string            `json:"metadata"  validate:""`
}

/////// response
type Response struct {
	StatusCode int   `json:"statuscode"  validate:"required"`
	Message    Frame `json:"message"  validate:"required"`
}
type GenericResponse struct {
	StateCode int    `json:"stateCode"  validate:"required"`
	Message   string `json:"message"  validate:"required"`
	Details   string `json:"details"  validate:"required"`
}
type SuccessGetDeviceResponse struct {
	StateCode int          `json:"stateCode"  validate:"required"`
	Message   string       `json:"message"  validate:"required"`
	Details   DeviceCreate `json:"details"  validate:"required"`
}
type SuccessGetRegistryResponse struct {
	StateCode int            `json:"stateCode"  validate:"required"`
	Message   string         `json:"message"  validate:"required"`
	Details   RegistryCreate `json:"details"  validate:"required"`
}
type SuccessGetDevicesResponse struct {
	StateCode int                    `json:"stateCode"  validate:"required"`
	Message   string                 `json:"message"  validate:"required"`
	Details   GetDevicesResultStruct `json:"details"  validate:"required"`
}
type SuccessGetRegistriesResponse struct {
	StateCode int                 `json:"stateCode"  validate:"required"`
	Message   string              `json:"message"  validate:"required"`
	Details   GetRegistriesResult `json:"details"  validate:"required"`
}
