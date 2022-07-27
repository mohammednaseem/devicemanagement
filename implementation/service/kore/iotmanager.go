package kore

import (
	"context"

	"github.com/gcp-iot/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type registryIotService struct {
	connectionString string
	Collection
	Client
}
type deviceIotService struct {
	connectionString string
	Collection
	Client
}

func NewRegistryService(conn string) model.IRegistryService {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	usersCollection := client.Database("testing").Collection("users")
	return &registryIotService{
		connectionString: conn,
	}
}
func NewDeviceService(conn string) model.IDeviceService {
	return &deviceIotService{
		connectionString: conn,
	}
}
