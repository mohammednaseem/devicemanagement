package usecase

import (
	"time"

	"github.com/gcp-iot/model"
)

type registryUsecase struct {
	registryService model.IRegistryrUsecase
	contextTimeout  time.Duration
}

func NewIoTUsecase(r model.IRegistryService, timeout time.Duration) model.IRegistryrUsecase {
	return &registryUsecase{
		registryService: r,
		contextTimeout:  timeout,
	}
}
