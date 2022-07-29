package http

import (
	"net/http"

	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func (r *registrytHandler) NewDevice(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(model.DeviceCreate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	req.Name = req.Parent + "/devices/" + req.Id
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.CreateDevice(ctx, *req)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) UpdateDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceUpdate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.UpdateMask = c.QueryParam("updateMask")
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	req.Name = req.Parent
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.UpdateDevice(ctx, *req)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) DeleteDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.DeleteDevice(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) GetDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.GetDevice(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
