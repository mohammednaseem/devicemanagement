package usecase

import (
	"context"
	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *registryUsecase) CreateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	c, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.CreateRegistry(c, registry)
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return dr, err

	}
	return dr, nil
}
func (i *registryUsecase) UpdateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	c, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.UpdateRegistry(c, registry)
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return dr, err

	}
	return dr, nil
}
func (i *registryUsecase) DeleteRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	c, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.DeleteRegistry(c, registry)
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return dr, err

	}
	return dr, nil
}
