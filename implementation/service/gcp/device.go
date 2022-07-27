package gcp

import (
	"context"
	"fmt"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	cloudiot "google.golang.org/api/cloudiot/v1"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (*deviceIotService) CreateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	var device cloudiot.Device

	// If no credentials are passed in, create an unauth device.
	if dev.PublicKeyFormat == "UNAUTH" {
		device = cloudiot.Device{
			Id: dev.DeviceID,
		}
	} else {
		device = cloudiot.Device{
			Id: dev.DeviceID,
			Credentials: []*cloudiot.DeviceCredential{
				{
					PublicKey: &cloudiot.PublicKeyCredential{
						Format: dev.PublicKeyFormat,
						Key:    dev.KeyBytes,
					},
				},
			},
		}
	}

	parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", dev.ProjectID, dev.Region, dev.RegistryID)
	_, err = client.Projects.Locations.Registries.Devices.Create(parent, &device).Do()
	if err != nil {

		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg(fmt.Sprintf("Successfully created a device with %s public key: %s", dev.PublicKeyFormat, dev.DeviceID))

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}

func (*deviceIotService) UpdateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{Message: err.Error()}
		return dr, err
	}

	path := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", dev.ProjectID, dev.Region, dev.RegistryID, dev.DeviceID)
	device, err := client.Projects.Locations.Registries.Devices.Get(path).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	device.Id = ""
	device.NumId = 0
	// If no credentials are passed in, create an unauth device.
	_, err = client.Projects.Locations.Registries.Devices.Patch(path, device).UpdateMask("blocked").Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg(fmt.Sprintf("Successfully Updated a device with %s public key: %s", dev.PublicKeyFormat, dev.DeviceID))

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (*deviceIotService) DeleteDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	path := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", dev.ProjectID, dev.Region, dev.RegistryID, dev.DeviceID)
	_, err = client.Projects.Locations.Registries.Devices.Delete(path).Do()
	if err != nil {
		//log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Deleted device: \n" + dev.DeviceID)

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
