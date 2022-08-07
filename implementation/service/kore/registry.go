package kore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/RacoWireless/iot-gw-thing-management/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/cloudiot/v1"
)

func CreateRegPublish(topicId string, dev model.RegistryCreate) error {

	PubStruct := model.PublishRegistryCreate{Operation: "POST", Entity: "Registry", Data: dev, Path: "registry/" + dev.Parent}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (r *registryIotService) CreateRegistry(_ context.Context, registry model.RegistryCreate) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id != "" {
		log.Error().Msg("Registry Already Exists")
		dr = model.FrameResponse(409, "Registry Already Exists", "")
		return dr, err
	}
	registry.CreatedOn = time.Now().String()
	insertOneResult, err := insertOne(r.ctx, r.client, r.database, r.collection, registry)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}
	log.Info().Msg("Result of InsertOne")
	log.Info().Msg((insertOneResult.InsertedID).(primitive.ObjectID).String())
	if r.Publish {
		err = CreateRegPublish(r.pubTopic, registry)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(201, "Success", "")
	return dr, err
}
func UpdateRegPublish(topicId string, dev model.RegistryUpdate) error {

	PubStruct := model.PublishRegistryUpdate{Operation: "PATCH", Entity: "Registry", Data: dev, Path: "registry/" + dev.Parent + "?updateMask=" + dev.UpdateMask}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}
func (r *registryIotService) UpdateRegistry(_ context.Context, registry model.RegistryUpdate) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	if registry.MqttConfig.MqttEnabledState != "" && strings.Contains(registry.UpdateMask, "mqtt_config") {
		queryResult.MqttConfig.MqttEnabledState = registry.MqttConfig.MqttEnabledState
	}
	if registry.HttpConfig.HttpEnabledState != "" && strings.Contains(registry.UpdateMask, "http_config") {
		queryResult.HttpConfig.HttpEnabledState = registry.HttpConfig.HttpEnabledState
	}
	if len(registry.EventNotificationConfigs) > 0 && strings.Contains(registry.UpdateMask, "event_notification_configs") {
		queryResult.EventNotificationConfigs = registry.EventNotificationConfigs
	}
	if registry.StateNotificationConfig != nil && strings.Contains(registry.UpdateMask, "state_notification_config") {
		queryResult.StateNotificationConfig = registry.StateNotificationConfig
	}
	if registry.LogLevel != "" && strings.Contains(registry.UpdateMask, "log_level") {
		queryResult.LogLevel = registry.LogLevel
	}
	if registry.Credentials != nil && strings.Contains(registry.UpdateMask, "credentials") {
		queryResult.Credentials = registry.Credentials
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "mqttconfig", Value: queryResult.MqttConfig},
		}}, {Key: "$set", Value: bson.D{
			{Key: "httpconfig", Value: queryResult.HttpConfig},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "eventnotificationconfigs", Value: queryResult.EventNotificationConfigs},
		}}, {Key: "$set", Value: bson.D{
			{Key: "statenotificationconfig", Value: queryResult.StateNotificationConfig},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "loglevel", Value: queryResult.LogLevel},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: queryResult.Credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(r.ctx, r.client, r.database, r.collection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if r.Publish {
		err = UpdateRegPublish(r.pubTopic, registry)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(200, "Success", "")
	return dr, err
}
func DeleteRegPublish(topicId string, dev model.RegistryDelete) error {

	PubStruct := model.PublishRegistryDelete{Operation: "DELETE", Entity: "Registry", Data: dev, Path: "registry/" + dev.Parent}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}
func (r *registryIotService) DeleteRegistry(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	Ping(r.ctx, r.client)
	filter := bson.D{{Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Parent}}}, {Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}}}
	var dr model.Response
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	if queryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(200, "Success", "")
		return dr, err
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "decomissioned", Value: true},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(r.ctx, r.client, r.database, r.collection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("Delete single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if r.Publish {
		err = DeleteRegPublish(r.pubTopic, registry)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(200, "Success", "")
	return dr, err
}
func (r *registryIotService) GetRegistry(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Parent}}}, {Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// Returns result of deletion and error
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	if queryResult.Id == "" {
		dr := model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	// print the count of affected documents
	log.Info().Msg("Got Details For Registry " + queryResult.Id)
	dr := model.FrameResponse(200, "Success", queryResult)
	return dr, err
}
func (r *registryIotService) GetRegistriesRegion(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "parent", Value: bson.D{{Key: "$eq", Value: registry.Parent}}}, {Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// Returns result of deletion and error
	//var queryResult model.RegistryCreate
	cursor, err := query(r.ctx, r.client, r.database, r.collection, filter)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	var results []model.RegistryCreate

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(r.ctx, &results); err != nil {

		// handle the error
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}
	if results == nil {
		dr := model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}

	// print the count of affected documents
	log.Info().Msg("Got Details For Registries ")
	dr := model.FrameResponse(200, "Success", model.GetRegistriesResult{DeviceRegistries: results})
	return dr, err
}
func (r *registryIotService) GetRegistries(_ context.Context, registry model.RegistryDelete) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "project", Value: bson.D{{Key: "$eq", Value: registry.Project}}}, {Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// Returns result of deletion and error
	//var queryResult model.RegistryCreate
	cursor, err := query(r.ctx, r.client, r.database, r.collection, filter)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	var results []model.RegistryCreate

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(r.ctx, &results); err != nil {

		// handle the error
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}
	if results == nil {
		dr := model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	// print the count of affected documents
	log.Info().Msg("Got Details For Registries ")
	dr := model.FrameResponse(200, "Success", model.GetRegistriesResult{DeviceRegistries: results})
	return dr, err
}
func AddRegCertificatePublish(topicId string, dev model.AddRegistryCert) error {

	PubStruct := model.PublishRegistryAddCert{Operation: "POST", Entity: "Registry", Data: dev, Path: "registry/" + dev.Parent}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}
func (r *registryIotService) AddCertificate(_ context.Context, registry model.AddRegistryCert) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	queryResult.Credentials = append(queryResult.Credentials, &registry.Credentials)

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: queryResult.Credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(r.ctx, r.client, r.database, r.collection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if r.Publish {
		err = AddRegCertificatePublish(r.pubTopic, registry)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(201, "Success", "")
	return dr, err
}
func DelRegCertificatePublish(topicId string, dev model.AddRegistryCert) error {

	PubStruct := model.PublishRegistryAddCert{Operation: "DELETE", Entity: "Registry", Data: dev, Path: "registry/" + dev.Parent}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}
func (r *registryIotService) DeleteCertificate(_ context.Context, registry model.AddRegistryCert) (model.Response, error) {
	Ping(r.ctx, r.client)
	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.RegistryCreate
	err := queryOne(r.ctx, r.client, r.database, r.collection, filter).Decode(&queryResult)
	var dr model.Response
	if queryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: registry.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: registry.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var certificate = registry.Credentials
	var credentials []cloudiot.RegistryCredential
	for _, cert := range queryResult.Credentials {
		if cert.PublicKeyCertificate.Certificate != certificate.PublicKeyCertificate.Certificate {
			credentials = append(credentials, *cert)
		}

	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(r.ctx, r.client, r.database, r.collection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if r.Publish {
		err = DelRegCertificatePublish(r.pubTopic, registry)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(200, "Success", "")
	return dr, err
}
