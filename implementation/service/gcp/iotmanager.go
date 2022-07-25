package gcp

import (
	"github.com/gcp-iot/model"
)

type registryIotService struct {
	connectionString string
}

func NewRegistryService(conn string) model.IRegistryService {
	return &registryIotService{
		connectionString: conn,
	}
}
