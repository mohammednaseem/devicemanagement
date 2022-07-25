package model

/////// request
type Request struct {
	ProjectID  string `json:"projectid" validate:"required"`
	Region     string `json:"region" validate:"required"`
	RegistryID string `json:"registryid" validate:"required"`
	TopicName  string `json:"topicname"  validate:"required"`
}

/////// response
type Response struct {
	StatusCode int    `json:"messgae"  validate:"required"`
	Message    string `json:"messgae"  validate:"required"`
}
