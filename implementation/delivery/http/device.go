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
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	if err := c.Validate(req); err != nil {
		return err
	}
	var reg model.RequestDevice = *req
	mResponse, err := r.dUsecase.CreateDevice(ctx, model.Device(reg))
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
	req.UpdateMask = c.QueryParam("updateMask")
	req.Parent = req.Name
	if err := c.Validate(req); err != nil {
		return err
	}
	var reg model.RequestDevice = *req
	mResponse, err := r.dUsecase.UpdateDevice(ctx, model.Device(reg))
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
	req.Parent = req.Name
	if err := c.Validate(req); err != nil {
		return err
	}
	var reg model.RequestDevice = *req
	mResponse, err := r.dUsecase.DeleteDevice(ctx, model.Device(reg))

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
