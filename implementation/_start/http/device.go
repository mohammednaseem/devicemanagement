package http

import (
	"net/http"

	"github.com/RacoWireless/iot-gw-thing-management/model"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

// CreateDevice godoc
// @Summary      Create Device
// @Description  create a device under a registry
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId}/devices [post]
func (r *registrytHandler) NewDevice(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(model.DeviceCreate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/devices"
	req.Name = req.Parent + "/" + req.Id
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Registry = c.Param("parent6")
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

// UpdateDevice godoc
// @Summary      Update Device
// @Description  update device under a registry
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Param        devId  path      string  true  "Device ID"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId}/devices/{devId} [patch]
func (r *registrytHandler) UpdateDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceUpdate)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.UpdateMask = c.QueryParam("updateMask")
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	req.Name = req.Parent
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Registry = c.Param("parent6")
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

// DeleteDevice godoc
// @Summary      Delete Device
// @Description  delete a device under a registry
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Param        devId  path      string  true  "Device ID"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId}/devices/{devId} [delete]
func (r *registrytHandler) DeleteDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = c.Param("parent8")
	req.Registry = c.Param("parent6")
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

// GetDevice godoc
// @Summary      Get Device
// @Description  Get a device under a registry
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Param        devId  path      string  true  "Device ID"
// @Success      200  {object}  model.Frame
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId}/devices/{devId} [get]
func (r *registrytHandler) GetDevice(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = c.Param("parent8")
	req.Registry = c.Param("parent6")
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

// GetDevices godoc
// @Summary      Get Devices
// @Description  Get all devices under a registry
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Success      200  {object}  model.GetDevicesResultStruct
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId}/devices [get]
func (r *registrytHandler) GetDevices(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.DeviceDelete)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Parent = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/devices"
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Registry = c.Param("parent6")
	req.Id = "ALL"
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.GetDevices(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}

// GetDevices godoc
// @Summary      Get Devices
// @Description  Get all devices under a registry
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project Id"
// @Param        region  path      string  true  "Region"
// @Param        registryId  path      string  true  "Registry ID"
// @Success      200  {object}  model.GetDevicesResultStruct
// @Failure      400  {object}  model.Frame
// @Failure      404  {object}  model.Frame
// @Failure      500  {object}  model.Frame
// @Router       /device/projects/{projectId}/locations/{region}/registries/{registryId}/devices [get]
func (r *registrytHandler) AddDevCertificate(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.AddDeviceCert)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Name = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	req.Parent = req.Parent + "/certificate"
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = c.Param("parent8")
	req.Registry = c.Param("parent6")
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.AddCertificate(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
func (r *registrytHandler) DelDevCertificate(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.AddDeviceCert)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("")
		r := model.FrameResponse(400, "Invalid Json Received", err.Error())
		return c.JSON(http.StatusBadRequest, r)
	}
	req.Name = c.Param("parent1") + "/" + c.Param("parent2") + "/" + c.Param("parent3") + "/" + c.Param("parent4") + "/" + c.Param("parent5") + "/" + c.Param("parent6") + "/" + c.Param("parent7") + "/" + c.Param("parent8")
	req.Parent = req.Parent + "/certificate"
	req.Project = c.Param("parent2")
	req.Region = c.Param("parent4")
	req.Id = c.Param("parent8")
	req.Registry = c.Param("parent6")
	if err := c.Validate(req); err != nil {
		return err
	}
	mResponse, err := r.dUsecase.DeleteCertificate(ctx, *req)

	if mResponse.StatusCode != 200 {
		log.Error().Err(err).Msg("")
		return c.JSON(mResponse.StatusCode, mResponse.Message)
	}
	return c.JSON(http.StatusOK, mResponse.Message)
}
