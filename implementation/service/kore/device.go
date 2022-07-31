package kore

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (d *deviceIotService) CreateDevice(_ context.Context, dev model.DeviceCreate) (model.Response, error) {
	ping(d.ctx, d.client)
	var rfilter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Registry}}},
		{Key: "region", Value: bson.D{{Key: "$eq", Value: dev.Region}}},
		{Key: "project", Value: bson.D{{Key: "$eq", Value: dev.Project}}},
	}
	var rqueryResult model.RegistryCreate
	var dr model.Response
	err := queryOne(d.ctx, d.client, d.database, d.rcollection, rfilter).Decode(&rqueryResult)
	if rqueryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.Response{StatusCode: 404, Message: "Registry Not Found"}
		return dr, err
	}
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
	}
	var queryResult model.DeviceCreate
	err = queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)

	if queryResult.Id != "" {
		log.Error().Msg("Device Already Exists")
		dr = model.Response{StatusCode: 409, Message: "Already Exists"}
		return dr, err
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(999999999999999999))
	if err != nil {
		log.Error().Msg("Random Generator Failed")
		dr = model.Response{StatusCode: 500, Message: "Internal Server Error"}
		return dr, err
	}
	randNum := nBig.Int64()
	dev.NumId = fmt.Sprintf("%d", randNum)
	insertOneResult, err := insertOne(d.ctx, d.client, d.database, d.dcollection, dev)
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
	err := queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
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
	updateResult, err := UpdateOne(d.ctx, d.client, d.database, d.dcollection, filter, update)
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
func (d *deviceIotService) DeleteDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
	}

	// Returns result of deletion and error
	result, err := deleteOne(d.ctx, d.client, d.database, d.dcollection, filter)
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
func (d *deviceIotService) GetDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
	}

	// Returns result of deletion and error
	var queryResult model.DeviceCreate
	err := queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	// print the count of affected documents
	if queryResult.Id == "" {
		dr := model.Response{StatusCode: 404, Message: "Not Result Found"}
		return dr, err
	}
	log.Info().Msg("Got Details For Device" + queryResult.Id)
	dr := model.Response{StatusCode: 200, Message: queryResult}
	return dr, err
}
func (d *deviceIotService) GetDevices(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "parent", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
	}

	// Returns result of deletion and error
	cursor, err := query(d.ctx, d.client, d.database, d.dcollection, filter)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	var results []model.DeviceCreate

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(d.ctx, &results); err != nil {

		// handle the error
		log.Error().Err(err).Msg("")
		dr := model.Response{StatusCode: 500, Message: err.Error()}
		return dr, err
	}
	if results == nil {
		dr := model.Response{StatusCode: 404, Message: "Not Result Found"}
		return dr, err
	}
	type resultNode struct {
		Id    string `json:"id" validate:"required"`
		NumID string `json:"numId" validate:"required"`
	}
	type resultStruct struct {
		Devices []resultNode `json:"devices" validate:"required"`
	}
	var result resultStruct
	for _, element := range results {
		node := resultNode{Id: element.Id, NumID: element.NumId}
		result.Devices = append(result.Devices, node)
	}

	// print the count of affected documents
	log.Info().Msg("Got Details For Devices ")
	dr := model.Response{StatusCode: 200, Message: result}
	return dr, err
}
