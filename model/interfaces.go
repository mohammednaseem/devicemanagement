package model

import "context"

//registry usecase
type IRegistryrUsecase interface {
	CreateRegistry(ctx context.Context, registry Registry) (Response, error)
	//update
	//delete
}

//registry service
type IRegistryService interface {
	CreateRegistry(ctx context.Context, registry Registry) (Response, error)
	//update
	//delete
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
