package gcp

import (
	"context"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2/google"
	cloudiot "google.golang.org/api/cloudiot/v1"
	"google.golang.org/api/option"
)

// getClient returns a client based on the environment variable GOOGLE_APPLICATION_CREDENTIALS
func getClient() (*cloudiot.Service, error) {
	// Authorize the client using Application Default Credentials.
	// See https://g.co/dv/identity/protocols/application-default-credentials
	ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, cloudiot.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	client, err := cloudiot.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (*registryIotService) CreateRegistry(_ context.Context, registry model.RegistryCreate) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	var devRegistry cloudiot.DeviceRegistry

	devRegistry.Id = registry.Id
	devRegistry.EventNotificationConfigs = registry.EventNotificationConfigs
	devRegistry.StateNotificationConfig = registry.StateNotificationConfig
	devRegistry.HttpConfig = &registry.HttpConfig
	devRegistry.MqttConfig = &registry.MqttConfig
	devRegistry.Credentials = registry.Credentials
	devRegistry.LogLevel = registry.LogLevel

	response, err := client.Projects.Locations.Registries.Create(registry.Parent, &devRegistry).Do()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Created registry:")
	log.Info().Msg(response.Id)
	log.Info().Msg(response.HttpConfig.HttpEnabledState)
	log.Info().Msg(response.MqttConfig.MqttEnabledState)
	log.Info().Msg(response.Name)

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}

func (*registryIotService) UpdateRegistry(_ context.Context, registry model.RegistryUpdate) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{Message: err.Error()}
		return dr, err
	}

	devRegistry, err := client.Projects.Locations.Registries.Get(registry.Parent).Do()
	if err != nil {
		dr := model.Response{Message: err.Error()}
		return dr, err
	}
	devRegistry.EventNotificationConfigs = registry.EventNotificationConfigs
	devRegistry.StateNotificationConfig = registry.StateNotificationConfig
	devRegistry.HttpConfig = &registry.HttpConfig
	devRegistry.MqttConfig = &registry.MqttConfig
	devRegistry.Id = ""
	response, err := client.Projects.Locations.Registries.Patch(registry.Parent, devRegistry).UpdateMask(registry.UpdateMask).Do()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Updated registry:")
	log.Info().Msg(response.Id)
	log.Info().Msg(response.HttpConfig.HttpEnabledState)
	log.Info().Msg(response.MqttConfig.MqttEnabledState)
	log.Info().Msg(response.Name)

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (*registryIotService) DeleteRegistry(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	_, err = client.Projects.Locations.Registries.Get(registry.Parent).Do()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	_, err = client.Projects.Locations.Registries.Delete(registry.Parent).Do()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Deleted registry:")

	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (*registryIotService) GetRegistry(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	reg, err := client.Projects.Locations.Registries.Get(registry.Parent).Do()
	if err != nil {
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	log.Info().Msg("Got registry:")

	dr := model.Response{StatusCode: 200, Message: reg}
	return dr, err
}
