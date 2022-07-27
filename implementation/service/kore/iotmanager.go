package kore

import (
	"context"

	"github.com/gcp-iot/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// query is user defined method used to query MongoDB,
// that accepts mongo.client,context, database name,
// collection name, a query and field.

//  database name and collection name is of type
// string. query is of type interface.
// field is of type interface, which limits
// the field being returned.

// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
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
