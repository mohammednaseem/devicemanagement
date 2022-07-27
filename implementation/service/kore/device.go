package kore

import (
	"context"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Error().Msg("Connection Unsuccessful")
		return err
	}
	log.Info().Msg("connected successfully")
	return nil
}

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (d *deviceIotService) CreateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	ping(d.client, d.ctx)
	var err error
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}

func (d *deviceIotService) UpdateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	ping(d.client, d.ctx)
	var err error
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (d *deviceIotService) DeleteDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	ping(d.client, d.ctx)
	var err error
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
