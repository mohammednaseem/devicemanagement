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
func (d *deviceIotService) CreateDevice(_ context.Context, dev model.DeviceCreate) (model.Response, error) {
	ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
	}
	var queryResult model.DeviceCreate
	err := queryOne(d.ctx, d.client, d.database, d.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id != "" {
		log.Error().Msg("Device Already Exists")
		dr = model.Response{StatusCode: 409, Message: "Already Exists"}
		return dr, err
	}
	insertOneResult, err := insertOne(d.ctx, d.client, d.database, d.collection, dev)
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

func (d *deviceIotService) UpdateDevice(_ context.Context, dev model.DeviceUpdate) (model.Response, error) {
	ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
	}
	var queryResult model.DeviceCreate
	err := queryOne(d.ctx, d.client, d.database, d.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.Response{StatusCode: 404, Message: "Not Found"}
		return dr, err
	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "blocked", Value: dev.Blocked},
		}}, {Key: "$set", Value: bson.D{
			{Key: "metadata", Value: dev.Metadata},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: dev.Credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(d.ctx, d.client, d.database, d.collection, filter, update)
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
func (d *deviceIotService) DeleteDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
	}

	// Returns result of deletion and error
	result, err := deleteOne(d.ctx, d.client, d.database, d.collection, filter)
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
