package kore

import (
	"context"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type registryIotService struct {
	client     *mongo.Client
	collection string
	database   string
	ctx        context.Context
}
type deviceIotService struct {
	client     *mongo.Client
	collection string
	database   string
	ctx        context.Context
}

func NewRegistryService(conn *mongo.Client, collection string, database string, ctx context.Context) model.IRegistryService {
	return &registryIotService{
		client:     conn,
		collection: collection,
		database:   database,
		ctx:        ctx,
	}
}
func NewDeviceService(conn *mongo.Client, collection string, database string, ctx context.Context) model.IDeviceService {
	return &deviceIotService{
		client:     conn,
		collection: collection,
		database:   database,
		ctx:        ctx,
	}
}

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

// insertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

// query is user defined method used to query MongoDB,
// that accepts mongo.client,context, database name,
// collection name, a query and field.

//  database name and collection name is of type
// string. query is of type interface.
// field is of type interface, which limits
// the field being returned.

// query method returns a cursor and error.
func queryOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.SingleResult) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result = collection.FindOne(ctx, query)
	return
}
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the database and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the
	// filter will get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

// deleteOne is a user defined function that delete,
// a single document from the collection.
// Returns DeleteResult and an  error if any.
func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document  from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}
