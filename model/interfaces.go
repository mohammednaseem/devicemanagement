package model

import "context"

//registry usecase
type IRegistryrUsecase interface {
	CreateRegistry(ctx context.Context, registry RegistryCreate) (Response, error)
	UpdateRegistry(ctx context.Context, registry RegistryUpdate) (Response, error)
	DeleteRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistriesRegion(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistries(ctx context.Context, registry RegistryDelete) (Response, error)
	AddCertificate(ctx context.Context, registry AddRegistryCert) (Response, error)
	DeleteCertificate(ctx context.Context, registry AddRegistryCert) (Response, error)
}

//registry service
type IRegistryService interface {
	CreateRegistry(ctx context.Context, registry RegistryCreate) (Response, error)
	UpdateRegistry(ctx context.Context, registry RegistryUpdate) (Response, error)
	DeleteRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistry(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistriesRegion(ctx context.Context, registry RegistryDelete) (Response, error)
	GetRegistries(ctx context.Context, registry RegistryDelete) (Response, error)
	AddCertificate(ctx context.Context, registry AddRegistryCert) (Response, error)
	DeleteCertificate(ctx context.Context, registry AddRegistryCert) (Response, error)
}

//device usecase
type IDevicerUsecase interface {
	CreateDevice(ctx context.Context, registry DeviceCreate) (Response, error)
	UpdateDevice(ctx context.Context, registry DeviceUpdate) (Response, error)
	DeleteDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevices(ctx context.Context, registry DeviceDelete) (Response, error)
	AddCertificate(ctx context.Context, registry AddDeviceCert) (Response, error)
	DeleteCertificate(ctx context.Context, registry AddDeviceCert) (Response, error)
}

//device service
type IDeviceService interface {
	CreateDevice(ctx context.Context, registry DeviceCreate) (Response, error)
	UpdateDevice(ctx context.Context, registry DeviceUpdate) (Response, error)
	DeleteDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevice(ctx context.Context, registry DeviceDelete) (Response, error)
	GetDevices(ctx context.Context, registry DeviceDelete) (Response, error)
	AddCertificate(ctx context.Context, registry AddDeviceCert) (Response, error)
	DeleteCertificate(ctx context.Context, registry AddDeviceCert) (Response, error)
}
