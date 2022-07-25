package http

import (
	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
)

type registrytHandler struct {
	rUsecase model.IRegistryrUsecase
}

func NewIoTtHandler(e *echo.Echo, registryUsecase model.IRegistryrUsecase) {
	RegistrytHandler := &registrytHandler{
		rUsecase: registryUsecase,
	}
	e.POST("/registry", RegistrytHandler.NewRegistry)
}
