package http

import (
	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (r *registrytHandler) NewRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.Request)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	reg := model.Registry{
		ProjectID:  req.ProjectID,
		Region:     req.Region,
		RegistryID: req.RegistryID,
		TopicName:  req.TopicName,
	}
	mResponse, err := r.rUsecase.CreateRegistry(ctx, reg)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(http.StatusInternalServerError, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
