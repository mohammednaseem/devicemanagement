package kore

import (
	"context"

	"github.com/gcp-iot/model"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (r *registryIotService) CreateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	ping(r.client, r.ctx)
	var err error
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}

func (r *registryIotService) UpdateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	ping(r.client, r.ctx)
	var err error
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (r *registryIotService) DeleteRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	ping(r.client, r.ctx)
	var err error
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
