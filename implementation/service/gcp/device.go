package gcp

import (
	"context"
	"fmt"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	cloudiot "google.golang.org/api/cloudiot/v1"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (*deviceIotService) CreateDevice(_ context.Context, dev model.DeviceCreate) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	var device cloudiot.Device

	device.Id = dev.Id
	device.Name = ""
	device.Blocked = dev.Blocked
	device.Credentials = dev.Credentials
	device.LogLevel = dev.LogLevel
	device.Metadata = dev.Metadata
	// If no credentials are passed in, create an unauth device.

	_, err = client.Projects.Locations.Registries.Devices.Create(dev.Parent, &device).Do()
	if err != nil {

		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Successfully created a device with  public key")

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}

func (*deviceIotService) UpdateDevice(_ context.Context, dev model.DeviceUpdate) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{Message: err.Error()}
		return dr, err
	}

	device, err := client.Projects.Locations.Registries.Devices.Get(dev.Parent).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	fmt.Print(device)
	device.Blocked = dev.Blocked
	device.Credentials = dev.Credentials
	device.Metadata = dev.Metadata
	device.Id = ""
	device.NumId = 0
	// If no credentials are passed in, create an unauth device.
	_, err = client.Projects.Locations.Registries.Devices.Patch(dev.Parent, device).UpdateMask(dev.UpdateMask).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Successfully Updated a device ")

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (*deviceIotService) DeleteDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	_, err = client.Projects.Locations.Registries.Devices.Delete(dev.Parent).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Deleted device: \n")

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (*deviceIotService) GetDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	device, err := client.Projects.Locations.Registries.Devices.Get(dev.Parent).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Got device: \n")

	dr := model.Response{StatusCode: 200, Message: device}
	return dr, err
}
func (*deviceIotService) GetDevices(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	device, err := client.Projects.Locations.Registries.Devices.List(dev.Parent).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Got device: \n")

	dr := model.Response{StatusCode: 200, Message: device}
	return dr, err
}
