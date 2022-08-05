package usecase

import (
	"context"

	"github.com/RacoWireless/iot-gw-thing-management/model"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *deviceUsecase) CreateDevice(ctx context.Context, dev model.DeviceCreate) (model.Response, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()
	dr, err := i.deviceService.CreateDevice(ctx, dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *deviceUsecase) UpdateDevice(ctx context.Context, dev model.DeviceUpdate) (model.Response, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.deviceService.UpdateDevice(ctx, dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *deviceUsecase) DeleteDevice(ctx context.Context, dev model.DeviceDelete) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.deviceService.DeleteDevice(ctx, dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *deviceUsecase) GetDevice(ctx context.Context, dev model.DeviceDelete) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.deviceService.GetDevice(ctx, dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
func (i *deviceUsecase) GetDevices(ctx context.Context, dev model.DeviceDelete) (model.Response, error) {
	var cancel context.CancelFunc
	_, cancel = context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	dr, err := i.deviceService.GetDevices(ctx, dev)
	if err != nil {

		return dr, err

	}
	return dr, nil
}
