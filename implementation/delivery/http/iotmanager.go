package http

import (
	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
)

type registrytHandler struct {
	rUsecase model.IRegistryrUsecase
	dUsecase model.IDevicerUsecase
}

func NewIoTtHandler(e *echo.Echo, registryUsecase model.IRegistryrUsecase, deviceUsecase model.IDevicerUsecase) {
	RegistrytHandler := &registrytHandler{
		rUsecase: registryUsecase,
		dUsecase: deviceUsecase,
	}
	e.POST("/registry", RegistrytHandler.NewRegistry)
	e.PATCH("/registry", RegistrytHandler.UpdateRegistry)
	e.DELETE("/registry", RegistrytHandler.DeleteRegistry)
	e.POST("/device", RegistrytHandler.NewDevice)
	e.PATCH("/device", RegistrytHandler.UpdateDevice)
	e.DELETE("/device", RegistrytHandler.DeleteDevice)
}
