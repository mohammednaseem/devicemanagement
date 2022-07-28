package kore

import (
	"context"
	"fmt"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (d *deviceIotService) CreateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	ping(d.client, d.ctx)
	var filter interface{} = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: dev.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: dev.ProjectID}}}, {Key: "deviceid", Value: bson.D{{Key: "$eq", Value: dev.DeviceID}}},
	}
	var queryResult model.Device
	err := queryOne(d.client, d.ctx, d.database, d.collection, filter).Decode(&queryResult)
	var dr model.Response
	if (queryResult != model.Device{}) {
		log.Error().Msg("Device Already Exists")
		dr = model.Response{StatusCode: 409, Message: "Already Exists"}
		return dr, err
	}
	insertOneResult, err := insertOne(d.client, d.ctx, d.database, d.collection, dev)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	log.Info().Msg("Result of InsertOne")
	log.Info().Msg((insertOneResult.InsertedID).(primitive.ObjectID).String())

	dr = model.Response{StatusCode: 200, Message: "Success"}

	return dr, err
}

func (d *deviceIotService) UpdateDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	ping(d.client, d.ctx)
	var filter interface{} = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: dev.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: dev.ProjectID}}}, {Key: "deviceid", Value: bson.D{{Key: "$eq", Value: dev.DeviceID}}},
	}
	var queryResult model.Device
	err := queryOne(d.client, d.ctx, d.database, d.collection, filter).Decode(&queryResult)
	var dr model.Response
	if (queryResult == model.Device{}) {
		log.Error().Msg("No Registry Found")
		dr = model.Response{StatusCode: 404, Message: "Not Found"}
		return dr, err
	}
	filter = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: dev.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: dev.ProjectID}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "publickeyformat", Value: dev.PublicKeyFormat},
		}}, {Key: "$set", Value: bson.D{
			{Key: "keybytes", Value: dev.KeyBytes},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(d.client, d.ctx, d.database, d.collection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	// print count of documents that affected
	fmt.Println("update single document")
	fmt.Println(updateResult.ModifiedCount)
	dr = model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (d *deviceIotService) DeleteDevice(ctx context.Context, dev model.Device) (model.Response, error) {
	ping(d.client, d.ctx)
	var filter interface{} = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: dev.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: dev.ProjectID}}}, {Key: "deviceid", Value: bson.D{{Key: "$eq", Value: dev.DeviceID}}},
	}

	// Returns result of deletion and error
	result, err := deleteOne(d.client, d.ctx, d.database, d.collection, filter)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	// print the count of affected documents
	log.Info().Msg("No.of rows affected by DeleteOne()")
	log.Info().Msg(fmt.Sprintf("%d", result.DeletedCount))
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
