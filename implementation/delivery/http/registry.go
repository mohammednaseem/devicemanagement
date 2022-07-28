package http

import (
	"net/http"

	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func (r *registrytHandler) NewRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RequestRegistry)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	reg := model.Registry{
		ProjectID:   req.ProjectID,
		Region:      req.Region,
		RegistryID:  req.RegistryID,
		TopicName:   req.TopicName,
		Certificate: req.Certificate,
	}
	mResponse, err := r.rUsecase.CreateRegistry(ctx, reg)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) UpdateRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RequestRegistry)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	reg := model.Registry{
		ProjectID:  req.ProjectID,
		Region:     req.Region,
		RegistryID: req.RegistryID,
		TopicName:  req.TopicName,
	}
	mResponse, err := r.rUsecase.UpdateRegistry(ctx, reg)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) DeleteRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RequestRegistry)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	reg := model.Registry{
		ProjectID:  req.ProjectID,
		Region:     req.Region,
		RegistryID: req.RegistryID,
		TopicName:  req.TopicName,
	}
	mResponse, err := r.rUsecase.DeleteRegistry(ctx, reg)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
