package usecase

import (
	"context"
	"log"

	"github.com/gcp-iot/model"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *registryUsecase) CreateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	c, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.registryService.CreateRegistry(c, registry)
	if err != nil {
		log.Fatal(err)
		return dr, err

	}
	return dr, nil
}
