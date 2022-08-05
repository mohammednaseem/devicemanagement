package http

import (
	"net/http"

	"github.com/RacoWireless/iot-gw-thing-management/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

// CreateRegistry godoc
// @Summary      Create Registry
// @Description  create a Registry
// @Tags         registry
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries [post]
func (r *registrytHandler) NewRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryCreate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/registries"
	req.Name = req.Parent + "/" + req.Id
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")

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

// UpdateRegistry godoc
// @Summary      Update Registry
// @Description  Update a Registry
// @Tags         registry
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId} [patch]
func (r *registrytHandler) UpdateRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryUpdate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.UpdateMask = c.QueryParam("updateMask")
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	req.Name = req.Parent
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
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

// DeleteRegistry godoc
// @Summary      Delete Registry
// @Description  Delete a Registry
// @Tags         registry
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId} [delete]
func (r *registrytHandler) DeleteRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = c.Param("parent6")
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

// GetRegistry godoc
// @Summary      Get Registry
// @Description  Get a Registry
// @Tags         registry
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Success      200  {object}  model.RegistryCreate
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId} [get]
func (r *registrytHandler) GetRegistry(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6")
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = c.Param("parent6")
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

// GetRegistriesRegion godoc
// @Summary      Get Registries With Region
// @Description  Get all Registries Under Region
// @Tags         registry
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Success      200  {object}  model.GetRegistriesResult
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries [get]
func (r *registrytHandler) GetRegistriesRegion(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/registries"
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = "ALL"
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.rUsecase.GetRegistriesRegion(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}

// GetRegistries godoc
// @Summary      Get Registries
// @Description  Get All Registries Under Project
// @Tags         registry
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Success      200  {object}  model.GetRegistriesResult
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/registries [get]
func (r *registrytHandler) GetRegistries(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.RegistryDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/registries"
	req.Project = c.Param("parent2")
	req.Id = "ALL"
	req.Region = "ALL"
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.rUsecase.GetRegistries(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
