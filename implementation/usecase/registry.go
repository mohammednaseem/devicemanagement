package usecase

import (
	"context"

	"github.com/gcp-iot/model"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *registryUsecase) CreateRegistry(ctx context.Context, registry model.RegistryCreate) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.CreateRegistry(ctx, registry)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *registryUsecase) UpdateRegistry(ctx context.Context, registry model.RegistryUpdate) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.UpdateRegistry(ctx, registry)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *registryUsecase) DeleteRegistry(ctx context.Context, registry model.RegistryDelete) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.DeleteRegistry(ctx, registry)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *registryUsecase) GetRegistry(ctx context.Context, registry model.RegistryDelete) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.GetRegistry(ctx, registry)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
