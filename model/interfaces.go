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
	CreateRegistry(ctx context.Context, registry Registry) (Response, error)
	UpdateRegistry(ctx context.Context, registry Registry) (Response, error)
	DeleteRegistry(ctx context.Context, registry Registry) (Response, error)
}

//device usecase
/*type IDeviceUsecase interface {
	CreateDevice(ctx context.Context, registry Registry) (Response, error)
	//update
	//delete
}

//device service
type IDeviceService interface {
	CreateDevice(ctx context.Context, registry Registry) (Response, error)
	//update
	//delete
}*/
