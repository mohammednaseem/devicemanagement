package gcp

import (
	"context"
	"fmt"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2/google"
	cloudiot "google.golang.org/api/cloudiot/v1"
	"google.golang.org/api/option"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (r *registryIotService) CreateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	devRegistry := cloudiot.DeviceRegistry{
		Id: registry.RegistryID,
		EventNotificationConfigs: []*cloudiot.EventNotificationConfig{
			{
				SubfolderMatches: "",
				PubsubTopicName:  registry.TopicName,
			},
		},
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", registry.ProjectID, registry.Region)
	response, err := client.Projects.Locations.Registries.Create(parent, &devRegistry).Do()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	log.Info().Msg("Created registry:")
	log.Info().Msg(response.Id)
	log.Info().Msg(response.HttpConfig.HttpEnabledState)
	log.Info().Msg(response.MqttConfig.MqttEnabledState)
	log.Info().Msg(response.Name)

	dr := model.Response{Message: "Success"}
	return dr, err
}

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

func (r *registryIotService) UpdateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", registry.ProjectID, registry.Region, registry.RegistryID)
	devRegistry, err := client.Projects.Locations.Registries.Get(parent).Do()
	devRegistry.EventNotificationConfigs = []*cloudiot.EventNotificationConfig{
		{
			PubsubTopicName: registry.TopicName,
		},
	}
	devRegistry.Id = ""
	response, err := client.Projects.Locations.Registries.Patch(parent, devRegistry).UpdateMask("event_notification_configs").Do()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	log.Info().Msg("Updated registry:")
	log.Info().Msg(response.Id)
	log.Info().Msg(response.HttpConfig.HttpEnabledState)
	log.Info().Msg(response.MqttConfig.MqttEnabledState)
	log.Info().Msg(response.Name)

	dr := model.Response{Message: "Success"}
	return dr, err
}
func (r *registryIotService) DeleteRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	client, err := getClient()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", registry.ProjectID, registry.Region, registry.RegistryID)
	_, err = client.Projects.Locations.Registries.Get(parent).Do()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	_, err = client.Projects.Locations.Registries.Delete(parent).Do()
	if err != nil {
		dr := model.Response{Message: "Error"}
		return dr, err
	}

	log.Info().Msg("Deleted registry:")

	dr := model.Response{Message: "Success"}
	return dr, err
}
