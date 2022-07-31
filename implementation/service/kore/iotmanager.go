package kore

import (
	"context"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type registryIotService struct {
	client     *mongo.Client
	collection string
	database   string
	ctx        context.Context
}
type deviceIotService struct {
	client      *mongo.Client
	rcollection string
	dcollection string
	database    string
	ctx         context.Context
}

func NewRegistryService(ctx context.Context, conn *mongo.Client, collection string, database string) model.IRegistryService {
	return &registryIotService{
		client:     conn,
		collection: collection,
		database:   database,
		ctx:        ctx,
	}
}
func NewDeviceService(ctx context.Context, conn *mongo.Client, dcollection string, rcollection string, database string) model.IDeviceService {
	return &deviceIotService{
		client:      conn,
		dcollection: dcollection, //device col
		rcollection: rcollection, //registry col
		database:    database,
		ctx:         ctx,
	}
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(ctx context.Context, client *mongo.Client) error {

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
func insertOne(ctx context.Context, client *mongo.Client, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

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
func query(ctx context.Context, client *mongo.Client, dataBase, col string, query interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query, options.Find().SetLimit(10))
	return
}

// query is user defined method used to query MongoDB,
// that accepts mongo.client,context, database name,
// collection name, a query and field.

//  database name and collection name is of type
// string. query is of type interface.
// field is of type interface, which limits
// the field being returned.

// query method returns a cursor and error.
func queryOne(ctx context.Context, client *mongo.Client, dataBase, col string, query interface{}) (result *mongo.SingleResult) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result = collection.FindOne(ctx, query)
	return
}
func UpdateOne(ctx context.Context, client *mongo.Client, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

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
func deleteOne(ctx context.Context, client *mongo.Client, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document  from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}
