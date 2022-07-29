package http

import (
	"net/http"

	"github.com/gcp-iot/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func (r *registrytHandler) NewRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryCreate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4")
	req.Name = req.Parent + "/registries/" + req.Id
	//req.Parent = "projects/my-iot-356305/locations/asia-east1"
	if err := c.Validate(req); err != nil {
		return err
	}
	// reg := model.Registry{
	// 	ProjectID:   req.ProjectID,
	// 	Region:      req.Region,
	// 	RegistryID:  req.RegistryID,
	// 	TopicName:   req.TopicName,
	// 	Certificate: req.Certificate,
	// }

	mResponse, err := r.rUsecase.CreateRegistry(ctx, *req)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) UpdateRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryUpdate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.UpdateMask = c.QueryParam("updateMask")
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	req.Name = req.Parent
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.rUsecase.UpdateRegistry(ctx, *req)
	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) DeleteRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.rUsecase.DeleteRegistry(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) GetRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.Response{Message: "Data not good"}
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.rUsecase.GetRegistry(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
