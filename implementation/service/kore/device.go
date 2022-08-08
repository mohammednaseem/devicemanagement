package kore

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/RacoWireless/iot-gw-thing-management/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
	"google.golang.org/api/cloudiot/v1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDevicePublish(topicId string, dev model.DeviceCreate) error {

	PubStruct := model.PublishDeviceCreate{Operation: "POST", Entity: "Device", Data: dev, Path: "device/" + dev.Parent}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (d *deviceIotService) CreateDevice(_ context.Context, dev model.DeviceCreate) (model.Response, error) {
	Ping(d.ctx, d.client)
	var rfilter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Registry}}},
		{Key: "region", Value: bson.D{{Key: "$eq", Value: dev.Region}}},
		{Key: "project", Value: bson.D{{Key: "$eq", Value: dev.Project}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var rqueryResult model.RegistryCreate
	var dr model.Response
	err := queryOne(d.ctx, d.client, d.database, d.rcollection, rfilter).Decode(&rqueryResult)
	if rqueryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}

	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.DeviceCreate
	err = queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)

	if queryResult.Id != "" {
		log.Error().Msg("Device Already Exists")
		dr = model.FrameResponse(409, "Device Already Exists", "")
		return dr, err
	}
	if len(rqueryResult.Credentials) > 0 {
		for _, cert := range dev.Credentials {
			for _, ca := range rqueryResult.Credentials {
				err = verifyCert([]byte(cert.PublicKey.Key), []byte(ca.PublicKeyCertificate.Certificate))
				if err == nil {
					break
				}
			}
			if err != nil {
				log.Error().Msg("Certificate Verification Failed")
				dr = model.FrameResponse(403, "Certificate Verification Failed", "")
				return dr, err
			}

		}

	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(999999999999999999))
	if err != nil {
		log.Error().Msg("Random Generator Failed")
		dr = model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}
	randNum := nBig.Int64()
	dev.NumId = fmt.Sprintf("%d", randNum)
	dev.CreatedOn = time.Now().String()
	insertOneResult, err := insertOne(d.ctx, d.client, d.database, d.dcollection, dev)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}
	log.Info().Msg("Result of InsertOne")
	log.Info().Msg((insertOneResult.InsertedID).(primitive.ObjectID).String())
	if d.Publish {
		err = CreateDevicePublish(d.pubTopic, dev)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(201, "Success", "")
	return dr, err
}
func UpdateDevicePublish(topicId string, dev model.DeviceUpdate) error {

	PubStruct := model.PublishDeviceUpdate{Operation: "PATCH", Entity: "Device", Data: dev, Path: "device/" + dev.Parent + "?updateMask=" + dev.UpdateMask}
	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)
	return err
}
func (d *deviceIotService) UpdateDevice(_ context.Context, dev model.DeviceUpdate) (model.Response, error) {
	Ping(d.ctx, d.client)

	var rfilter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Registry}}},
		{Key: "region", Value: bson.D{{Key: "$eq", Value: dev.Region}}},
		{Key: "project", Value: bson.D{{Key: "$eq", Value: dev.Project}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var rqueryResult model.RegistryCreate
	var dr model.Response
	err := queryOne(d.ctx, d.client, d.database, d.rcollection, rfilter).Decode(&rqueryResult)
	if rqueryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}

	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.DeviceCreate
	err = queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
	if queryResult.Id == "" {
		log.Error().Msg("No Device Found")
		dr = model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}
	if len(rqueryResult.Credentials) > 0 {
		var noCerts []string
		for _, cert := range dev.Credentials {
			for _, queryCert := range queryResult.Credentials {
				if cert.PublicKey.Key == queryCert.PublicKey.Key {
					noCerts = append(noCerts, cert.PublicKey.Key)
				}
			}
		}
		for _, cert := range dev.Credentials {
			if !slices.Contains(noCerts, cert.PublicKey.Key) {
				for _, ca := range rqueryResult.Credentials {
					err = verifyCert([]byte(cert.PublicKey.Key), []byte(ca.PublicKeyCertificate.Certificate))
					if err == nil {
						break
					}
				}
				if err != nil {
					log.Error().Msg("Certificate Verification Failed")
					dr = model.FrameResponse(403, "Certificate Verification Failed", "")
					return dr, err
				}
			}

		}

	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	if strings.Contains(dev.UpdateMask, "blocked") {
		queryResult.Blocked = dev.Blocked
	}
	if dev.Metadata != nil && strings.Contains(dev.UpdateMask, "metadata") {
		queryResult.Metadata = dev.Metadata
	}
	if len(dev.Credentials) > 0 && strings.Contains(dev.UpdateMask, "credentials") {
		queryResult.Credentials = dev.Credentials
	}
	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "blocked", Value: queryResult.Blocked},
		}}, {Key: "$set", Value: bson.D{
			{Key: "metadata", Value: queryResult.Metadata},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: queryResult.Credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(d.ctx, d.client, d.database, d.dcollection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if d.Publish {
		err = UpdateDevicePublish(d.pubTopic, dev)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(200, "Success", "")
	return dr, err
}
func DeleteDevicePublish(topicId string, dev model.DeviceDelete) error {

	PubStruct := model.PublishDeviceDelete{Operation: "DELETE", Entity: "Device", Data: dev, Path: "device/" + dev.Parent}

	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)

	return err
}

func (d *deviceIotService) DeleteDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	Ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.DeviceCreate
	var dr model.Response
	err := queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
	if queryResult.Id == "" {
		log.Error().Msg("No Device Found")
		dr = model.FrameResponse(200, "Device Deleted", "")
		return dr, err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "decomissioned", Value: true},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(d.ctx, d.client, d.database, d.dcollection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("Delete Single Device")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if d.Publish {
		err = DeleteDevicePublish(d.pubTopic, dev)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(200, "Success", "")
	return dr, err
}
func (d *deviceIotService) GetDevice(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	Ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// Returns result of deletion and error
	var queryResult model.DeviceCreate
	err := queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}
	// print the count of affected documents
	if queryResult.Id == "" {
		dr := model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}
	log.Info().Msg("Got Details For Device" + queryResult.Id)
	dr := model.FrameResponse(200, "Success", queryResult)
	return dr, err
}
func (d *deviceIotService) GetDevices(_ context.Context, dev model.DeviceDelete) (model.Response, error) {
	Ping(d.ctx, d.client)
	var filter interface{} = bson.D{
		{Key: "parent", Value: bson.D{{Key: "$eq", Value: dev.Parent}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// Returns result of deletion and error
	cursor, err := query(d.ctx, d.client, d.database, d.dcollection, filter)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}
	var results []model.DeviceCreate

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(d.ctx, &results); err != nil {

		// handle the error
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}
	if results == nil {
		dr := model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}

	var result model.GetDevicesResultStruct
	for _, element := range results {
		node := model.GetDevicesResultNode{Id: element.Id, NumID: element.NumId, Blocked: element.Blocked, LogLevel: element.LogLevel}
		result.Devices = append(result.Devices, node)
	}

	// print the count of affected documents
	log.Info().Msg("Got Details For Devices ")
	dr := model.FrameResponse(200, "Success", result)
	return dr, err
}
func AddDevCertificatePublish(topicId string, dev model.AddDeviceCert) error {

	PubStruct := model.PublishDeviceAddCert{Operation: "POST", Entity: "Device", Data: dev, Path: "device/" + dev.Parent}
	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)
	return err
}
func (d *deviceIotService) AddCertificate(_ context.Context, dev model.AddDeviceCert) (model.Response, error) {
	Ping(d.ctx, d.client)

	var rfilter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Registry}}},
		{Key: "region", Value: bson.D{{Key: "$eq", Value: dev.Region}}},
		{Key: "project", Value: bson.D{{Key: "$eq", Value: dev.Project}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var rqueryResult model.RegistryCreate
	var dr model.Response
	err := queryOne(d.ctx, d.client, d.database, d.rcollection, rfilter).Decode(&rqueryResult)
	if rqueryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}

	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.DeviceCreate
	err = queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
	if queryResult.Id == "" {
		log.Error().Msg("No Device Found")
		dr = model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}
	var certificate = dev.Credentials
	if len(rqueryResult.Credentials) > 0 {

		for _, ca := range rqueryResult.Credentials {
			err = verifyCert([]byte(certificate.PublicKey.Key), []byte(ca.PublicKeyCertificate.Certificate))
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Error().Msg("Certificate Verification Failed")
			dr = model.FrameResponse(403, "Certificate Verification Failed", "")
			return dr, err
		}
	}
	queryResult.Credentials = append(queryResult.Credentials, &certificate)
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: queryResult.Credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(d.ctx, d.client, d.database, d.dcollection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if d.Publish {
		err = AddDevCertificatePublish(d.pubTopic, dev)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(201, "Success", "")
	return dr, err
}
func DeleteDevCertificatePublish(topicId string, dev model.AddDeviceCert) error {

	PubStruct := model.PublishDeviceAddCert{Operation: "DELETE", Entity: "Device", Data: dev, Path: "device/" + dev.Parent}
	msg, err := json.Marshal(PubStruct)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = publish(dev.Project, topicId, msg)
	return err
}
func (d *deviceIotService) DeleteCertificate(_ context.Context, dev model.AddDeviceCert) (model.Response, error) {
	Ping(d.ctx, d.client)

	var rfilter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Registry}}},
		{Key: "region", Value: bson.D{{Key: "$eq", Value: dev.Region}}},
		{Key: "project", Value: bson.D{{Key: "$eq", Value: dev.Project}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var rqueryResult model.RegistryCreate
	var dr model.Response
	err := queryOne(d.ctx, d.client, d.database, d.rcollection, rfilter).Decode(&rqueryResult)
	if rqueryResult.Id == "" {
		log.Error().Msg("No Registry Found")
		dr = model.FrameResponse(404, "Registry Not Found", "")
		return dr, err
	}

	var filter interface{} = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}
	var queryResult model.DeviceCreate
	err = queryOne(d.ctx, d.client, d.database, d.dcollection, filter).Decode(&queryResult)
	if queryResult.Id == "" {
		log.Error().Msg("No Device Found")
		dr = model.FrameResponse(404, "Device Not Found", "")
		return dr, err
	}
	var certificate = dev.Credentials
	var credentials []cloudiot.DeviceCredential
	for _, cert := range queryResult.Credentials {
		if cert.PublicKey.Key != certificate.PublicKey.Key {
			credentials = append(credentials, *cert)
		}

	}
	filter = bson.D{
		{Key: "id", Value: bson.D{{Key: "$eq", Value: dev.Id}}}, {Key: "name", Value: bson.D{{Key: "$eq", Value: dev.Name}}},
		{Key: "decomissioned", Value: bson.D{{Key: "$eq", Value: false}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "credentials", Value: credentials},
		}},
	}

	// Returns result of updated document and a error.
	updateResult, err := UpdateOne(d.ctx, d.client, d.database, d.dcollection, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("")
		dr := model.FrameResponse(500, "Internal Server Error", err.Error())
		return dr, err
	}

	// print count of documents that affected
	log.Info().Msg("update single document")
	log.Info().Msg(fmt.Sprintf("%d", updateResult.ModifiedCount))
	if d.Publish {
		err = DeleteDevCertificatePublish(d.pubTopic, dev)
		if err != nil {
			dr := model.FrameResponse(500, "Internal Server Error", err.Error())
			return dr, err
		}
	}
	dr = model.FrameResponse(200, "Success", "")
	return dr, err
}
