package model

/////// request
type RequestRegistry struct {
	ProjectID   string `json:"projectid" validate:"required"`
	Region      string `json:"region" validate:"required"`
	RegistryID  string `json:"registryid" validate:"required"`
	TopicName   string `json:"topicname"  validate:"required"`
	Certificate string `json:"certificate"  validate:""`
}
type RequestDevice struct {
	ProjectID       string `json:"projectid" validate:"required"`
	Region          string `json:"region" validate:"required"`
	RegistryID      string `json:"registryid" validate:"required"`
	PublicKeyFormat string `json:"publickeyformat"  validate:"required"`
	KeyBytes        string `json:"keybytes"  validate:""`
	DeviceID        string `json:"deviceid"  validate:"required"`
}

/////// response
type Response struct {
	StatusCode int    `json:"messgae"  validate:"required"`
	Message    string `json:"messgae"  validate:"required"`
}
