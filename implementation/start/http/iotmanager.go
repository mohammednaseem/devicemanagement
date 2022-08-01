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
	e.GET("/registry/:parent1/:parent2/:parent3/:parent4/registries", RegistrytHandler.GetRegistries)
	e.POST("/registry/:parent1/:parent2/:parent3/:parent4/registries", RegistrytHandler.NewRegistry)
	e.PATCH("/registry/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6", RegistrytHandler.UpdateRegistry)
	e.DELETE("/registry/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6", RegistrytHandler.DeleteRegistry)
	e.GET("/registry/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6", RegistrytHandler.GetRegistry)
	e.POST("/device/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6/devices", RegistrytHandler.NewDevice)
	e.PATCH("/device/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6/:parent7/:parent8", RegistrytHandler.UpdateDevice)
	e.DELETE("/device/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6/:parent7/:parent8", RegistrytHandler.DeleteDevice)
	e.GET("/device/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6/:parent7/:parent8", RegistrytHandler.GetDevice)
	e.GET("/device/:parent1/:parent2/:parent3/:parent4/:parent5/:parent6/devices", RegistrytHandler.GetDevices)

}
