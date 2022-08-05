package usecase

import (
	"time"

	"github.com/RacoWireless/iot-gw-thing-management/model"
)

type registryUsecase struct {
	registryService model.IRegistryService
	contextTimeout  time.Duration
}
type deviceUsecase struct {
	deviceService  model.IDeviceService
	contextTimeout time.Duration
}

func NewIoTUsecase(r model.IRegistryService, timeout time.Duration) model.IRegistryrUsecase {
	return &registryUsecase{
		registryService: r,
		contextTimeout:  timeout,
	}
}
func NewDeviceUsecase(r model.IDeviceService, timeout time.Duration) model.IDevicerUsecase {
	return &deviceUsecase{
		deviceService:  r,
		contextTimeout: timeout,
	}
}
