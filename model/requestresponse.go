package model

import (
	cloudiot "google.golang.org/api/cloudiot/v1"
)

/////// request
type RequestRegistry struct {
	UpdateMask               string                              `json:"updatemask" validate:""`
	Parent                   string                              `json:"parent" validate:"required"`
	Id                       string                              `json:"id" validate:"required"`
	Name                     string                              `json:"name" validate:"required"`
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
	Name        string                       `json:"name" validate:"required"`
	Credentials []*cloudiot.DeviceCredential `json:"credentials" validate:"required"`
	LogLevel    string                       `json:"logLevel"  validate:""`
	Blocked     bool                         `json:"blocked"  validate:""`
	Metadata    map[string]string            `json:"metadata"  validate:""`
}

/////// response
type Response struct {
	StatusCode int         `json:"statuscode"  validate:"required"`
	Message    interface{} `json:"message"  validate:"required"`
}
