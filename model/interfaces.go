package model

import "context"

//registry usecase
type IRegistryrUsecase interface {
	CreateRegistry(ctx context.Context, registry RegistryCreate) (Response, error)
	UpdateRegistry(ctx context.Context, registry RegistryUpdate) (Response, error)
	DeleteRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistries(ctx context.Context, registry RegistryDelete) (Response, error)
}

//registry service
type IRegistryService interface {
	CreateRegistry(ctx context.Context, registry RegistryCreate) (Response, error)
	UpdateRegistry(ctx context.Context, registry RegistryUpdate) (Response, error)
	DeleteRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistries(ctx context.Context, registry RegistryDelete) (Response, error)
}

//device usecase
type IDevicerUsecase interface {
	CreateDevice(ctx context.Context, registry DeviceCreate) (Response, error)
	UpdateDevice(ctx context.Context, registry DeviceUpdate) (Response, error)
	DeleteDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevices(ctx context.Context, registry DeviceDelete) (Response, error)
}

//device service
type IDeviceService interface {
	CreateDevice(ctx context.Context, registry DeviceCreate) (Response, error)
	UpdateDevice(ctx context.Context, registry DeviceUpdate) (Response, error)
	DeleteDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevices(ctx context.Context, registry DeviceDelete) (Response, error)
}
