package model

type Registry struct {
	ProjectID  string `json:"projectid" validate:"required"`
	Region     string `json:"region" validate:"required"`
	RegistryID string `json:"registryid" validate:"required"`
	TopicName  string `json:"topicname"  validate:"required"`
}
