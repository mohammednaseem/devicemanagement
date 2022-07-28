package http

import (
	"net/http"

	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func (r *registrytHandler) NewDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RequestDevice)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	reg := model.Device{
		ProjectID:       req.ProjectID,
		Region:          req.Region,
		RegistryID:      req.RegistryID,
		PublicKeyFormat: req.PublicKeyFormat,
		KeyBytes:        req.KeyBytes,
		DeviceID:        req.DeviceID,
	}
	mResponse, err := r.dUsecase.CreateDevice(ctx, reg)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) UpdateDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RequestDevice)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	reg := model.Device{
		ProjectID:       req.ProjectID,
		Region:          req.Region,
		RegistryID:      req.RegistryID,
		PublicKeyFormat: req.PublicKeyFormat,
		KeyBytes:        req.KeyBytes,
		DeviceID:        req.DeviceID,
	}
	mResponse, err := r.dUsecase.UpdateDevice(ctx, reg)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) DeleteDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RequestDevice)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	reg := model.Device{
		ProjectID:       req.ProjectID,
		Region:          req.Region,
		RegistryID:      req.RegistryID,
		PublicKeyFormat: req.PublicKeyFormat,
		KeyBytes:        req.KeyBytes,
		DeviceID:        req.DeviceID,
	}
	mResponse, err := r.dUsecase.DeleteDevice(ctx, reg)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
