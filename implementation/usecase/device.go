package usecase

import (
	"context"

	"github.com/gcp-iot/model"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *deviceUsecase) CreateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()
	dr, err := i.deviceService.CreateDevice(dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *deviceUsecase) UpdateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.deviceService.UpdateDevice(dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *deviceUsecase) DeleteDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.deviceService.DeleteDevice(dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
