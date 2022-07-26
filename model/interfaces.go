package model

import "context"

//registry usecase
type IRegistryrUsecase interface {
	CreateRegistry(ctx context.Context, registry Registry) (Response, error)
	UpdateRegistry(ctx context.Context, registry Registry) (Response, error)
	DeleteRegistry(ctx context.Context, registry Registry) (Response, error)
}

//registry service
type IRegistryService interface {
	CreateRegistry(registry Registry) (Response, error)
	UpdateRegistry(registry Registry) (Response, error)
	DeleteRegistry(registry Registry) (Response, error)
}

//device usecase
type IDevicerUsecase interface {
	CreateDevice(ctx context.Context, registry Device) (Response, error)
	UpdateDevice(ctx context.Context, registry Device) (Response, error)
	DeleteDevice(ctx context.Context, registry Device) (Response, error)
}

//device service
type IDeviceService interface {
	CreateDevice(registry Device) (Response, error)
	UpdateDevice(registry Device) (Response, error)
	DeleteDevice(registry Device) (Response, error)
}
