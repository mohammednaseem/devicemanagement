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
func (r *registryIotService) CreateRegistry(_ context.Context, registry model.RegistryCreate) (model.Response, error) {
	ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
	}
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id != "" {
		log.Error().Msg("Registry Already Exists")
		dr = model.Response{StatusCode: 409, Message: "Already Exists"}
		return dr, err
	}
	insertOneResult, err := insertOne(r.ctx, r.client, r.database, r.collection, registry)
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

func (r *registryIotService) UpdateRegistry(_ context.Context, registry model.RegistryUpdate) (model.Response, error) {
	ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
	}
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.Response{StatusCode: 404, Message: "Not Found"}
		return dr, err
	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "mqttconfig", Value: registry.MqttConfig},
		}}, {Key: "$set", Value: bson.D{
			{Key: "httpconfig", Value: registry.HttpConfig},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: registry.Credentials},
		}}, {Key: "$set", Value: bson.D{
			{Key: "loglevel", Value: registry.LogLevel},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "eventnotificationconfigs", Value: registry.EventNotificationConfigs},
		}}, {Key: "$set", Value: bson.D{
			{Key: "statenotificationconfig", Value: registry.StateNotificationConfig},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(r.ctx, r.client, r.database, r.collection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	dr = model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
func (r *registryIotService) DeleteRegistry(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Parent}}},
	}

	// Returns result of deletion and error
	result, err := deleteOne(r.ctx, r.client, r.database, r.collection, filter)
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
func (r *registryIotService) GetRegistry(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Parent}}},
	}

	// Returns result of deletion and error
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	// print the count of affected documents
	log.Info().Msg("Got Details For Registry " + queryResult.Id)
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
