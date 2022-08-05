package model

import (
	cloudiot "google.golang.org/api/cloudiot/v1"
)

type RegistryCreate struct {
	Parent                   string                              `json:"parent" validate:"required"`
	Project                  string                              `json:"project" validate:"required"`
	Region                   string                              `json:"region" validate:"required"`
	Id                       string                              `json:"id" validate:"required"`
	Name                     string                              `json:"name" validate:"required"`
	EventNotificationConfigs []*cloudiot.EventNotificationConfig `json:"eventNotificationConfigs" validate:"required"`
	StateNotificationConfig  *cloudiot.StateNotificationConfig   `json:"stateNotificationConfig"  validate:""`
	MqttConfig               cloudiot.MqttConfig                 `json:"mqttConfig"  validate:""`
	HttpConfig               cloudiot.HttpConfig                 `json:"httpConfig"  validate:""`
	LogLevel                 string                              `json:"logLevel"  validate:""`
	CreatedOn                string                              `json:"createdOn"  validate:""`
	Credentials              []*cloudiot.RegistryCredential      `json:"credentials"  validate:""`
	Decomissioned            bool                                `json:"decomissioned"  validate:""`
}
type RegistryUpdate struct {
	UpdateMask               string                              `json:"updatemask" validate:"required"`
	Project                  string                              `json:"project" validate:"required"`
	Region                   string                              `json:"region" validate:"required"`
	Parent                   string                              `json:"parent" validate:"required"`
	Id                       string                              `json:"id" validate:"required"`
	Name                     string                              `json:"name" validate:"required"`
	EventNotificationConfigs []*cloudiot.EventNotificationConfig `json:"eventNotificationConfigs" validate:""`
	StateNotificationConfig  *cloudiot.StateNotificationConfig   `json:"stateNotificationConfig"  validate:""`
	MqttConfig               cloudiot.MqttConfig                 `json:"mqttConfig"  validate:""`
	HttpConfig               cloudiot.HttpConfig                 `json:"httpConfig"  validate:""`
	LogLevel                 string                              `json:"logLevel"  validate:""`
	CreatedOn                string                              `json:"createdOn"  validate:""`
	Credentials              []*cloudiot.RegistryCredential      `json:"credentials"  validate:""`
}
type RegistryDelete struct {
	Id      string `json:"id" validate:"required"`
	Parent  string `json:"parent" validate:"required"`
	Project string `json:"project" validate:"required"`
	Region  string `json:"region" validate:"required"`
}
type DeviceCreate struct {
	Project       string                       `json:"project" validate:"required"`
	Parent        string                       `json:"parent" validate:"required"`
	NumId         string                       `json:"numId" validate:""`
	Region        string                       `json:"region" validate:"required"`
	Registry      string                       `json:"registry" validate:"required"`
	Id            string                       `json:"id" validate:"required"`
	Name          string                       `json:"name" validate:"required"`
	Credentials   []*cloudiot.DeviceCredential `json:"credentials" validate:"required"`
	LogLevel      string                       `json:"logLevel"  validate:""`
	Blocked       bool                         `json:"blocked"  validate:""`
	Metadata      map[string]string            `json:"metadata"  validate:""`
	CreatedOn     string                       `json:"createdOn"  validate:""`
	Decomissioned bool                         `json:"decomissioned"  validate:""`
}
type DeviceUpdate struct {
	UpdateMask  string                       `json:"updatemask" validate:"required"`
	Project     string                       `json:"project" validate:"required"`
	Region      string                       `json:"region" validate:"required"`
	Registry    string                       `json:"registry" validate:"required"`
	Parent      string                       `json:"parent" validate:"required"`
	Id          string                       `json:"id" validate:"required"`
	Name        string                       `json:"name" validate:"required"`
	Credentials []*cloudiot.DeviceCredential `json:"credentials" validate:"required"`
	LogLevel    string                       `json:"logLevel"  validate:""`
	Blocked     bool                         `json:"blocked"  validate:""`
	Metadata    map[string]string            `json:"metadata"  validate:""`
	CreatedOn   string                       `json:"createdOn"  validate:""`
}
type DeviceDelete struct {
	Id       string `json:"id" validate:"required"`
	Parent   string `json:"parent" validate:"required"`
	Project  string `json:"project" validate:"required"`
	Region   string `json:"region" validate:"required"`
	Registry string `json:"registry" validate:"required"`
}
type GetDevicesResultNode struct {
	Id       string `json:"id" validate:"required"`
	NumID    string `json:"numId" validate:"required"`
	Blocked  bool   `json:"blocked" validate:"required"`
	LogLevel string `json:"loglevel" validate:"required"`
}
type GetDevicesResultStruct struct {
	Devices []GetDevicesResultNode `json:"devices" validate:"required"`
}
type GetRegistriesResult struct {
	DeviceRegistries []RegistryCreate `json:"deviceRegistries" validate:"required"`
}
