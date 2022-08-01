package kore

import (
	"context"

	"github.com/gcp-iot/model"
	"go.mongodb.org/mongo-driver/mongo"
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
	pubTopic    string
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
func NewDeviceService(ctx context.Context, conn *mongo.Client, dcollection string, rcollection string, database string, PubTopic string) model.IDeviceService {
	return &deviceIotService{
		client:      conn,
		dcollection: dcollection, //device col
		rcollection: rcollection, //registry col
		database:    database,
		ctx:         ctx,
		pubTopic:    PubTopic,
	}
}
