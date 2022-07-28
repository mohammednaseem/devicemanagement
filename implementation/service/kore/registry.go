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
func (r *registryIotService) CreateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	ping(r.client, r.ctx)
	var filter interface{} = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: registry.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: registry.ProjectID}}},
	}
	var queryResult model.Device
	err := queryOne(r.client, r.ctx, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if (queryResult != model.Device{}) {
		log.Error().Msg("Registry Already Exists")
		dr = model.Response{StatusCode: 409, Message: "Already Exists"}
		return dr, err
	}
	insertOneResult, err := insertOne(r.client, r.ctx, r.database, r.collection, registry)
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

func (r *registryIotService) UpdateRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	ping(r.client, r.ctx)
	var filter interface{} = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: registry.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: registry.ProjectID}}},
	}
	var queryResult model.Device
	err := queryOne(r.client, r.ctx, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if (queryResult == model.Device{}) {
		log.Error().Msg("No Registry Found")
		dr = model.Response{StatusCode: 404, Message: "Not Found"}
		return dr, err
	}
	filter = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: registry.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: registry.ProjectID}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "topicname", Value: registry.TopicName},
		}}, {Key: "$set", Value: bson.D{
			{Key: "certificate", Value: registry.Certificate},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(r.client, r.ctx, r.database, r.collection, filter, update)
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
func (r *registryIotService) DeleteRegistry(ctx context.Context, registry model.Registry) (model.Response, error) {
	ping(r.client, r.ctx)
	var filter interface{} = bson.D{
		{Key: "registryid", Value: bson.D{{Key: "$eq", Value: registry.RegistryID}}}, {Key: "projectid", Value: bson.D{{Key: "$eq", Value: registry.ProjectID}}},
	}

	// Returns result of deletion and error
	result, err := deleteOne(r.client, r.ctx, r.database, r.collection, filter)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	// print the count of affected documents
	log.Info().Msg("No.of rows affected by DeleteOne()")
	log.Info().Msg(fmt.Sprintf("%r", result.DeletedCount))
	dr := model.Response{StatusCode: 200, Message: "Success"}
	return dr, err
}
