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
	pubTopic   string
	ctx        context.Context
	Publish    bool
}
type deviceIotService struct {
	client      *mongo.Client
	rcollection string
	dcollection string
	database    string
	pubTopic    string
	ctx         context.Context
	Publish     bool
}

func NewRegistryService(ctx context.Context, conn *mongo.Client, collection string, database string, PubTopic string, Publish bool) model.IRegistryService {
	return &registryIotService{
		client:     conn,
		collection: collection,
		database:   database,
		ctx:        ctx,
		pubTopic:   PubTopic,
		Publish:    Publish,
	}
}
func NewDeviceService(ctx context.Context, conn *mongo.Client, dcollection string, rcollection string, database string, PubTopic string, Publish bool) model.IDeviceService {
	return &deviceIotService{
		client:      conn,
		dcollection: dcollection, //device col
		rcollection: rcollection, //registry col
		database:    database,
		ctx:         ctx,
		pubTopic:    PubTopic,
		Publish:     Publish,
	}
}
