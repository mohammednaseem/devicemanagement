package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	iotDelivery "github.com/gcp-iot/implementation/delivery/http"
	gcpService "github.com/gcp-iot/implementation/service/gcp"
	koreService "github.com/gcp-iot/implementation/service/kore"
	iotUsecase "github.com/gcp-iot/implementation/usecase"
	"github.com/gcp-iot/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/spf13/viper"

	Log "github.com/labstack/gommon/log"
	"github.com/rs/zerolog/log"
	lecho "github.com/ziflex/lecho"
)

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	log.Info().Msg("Closing Mongo Conection")
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns
// a mongo.Client, context.Context,
// context.CancelFunc and error.
// mongo.Client will be used for further database
// operation. context.Context will be used set
// deadlines for process. context.CancelFunc will
// be used to cancel context and resource
// associated with it.
func connect(uri string) (*mongo.Client, context.Context, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, err
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
func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	log.Info().Msg(`path: ` + path)
	viper.SetConfigType(`json`)
	viper.SetConfigName(`config`)
	viper.AddConfigPath(`./`)
	viper.AddConfigPath(`../`)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error().Err(err).Msg("config file not found")
		}
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Info().Msg("Service RUN on DEBUG mode")
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
func main() {

	log.Info().Msg("Go Time")

	flag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvPrefix(viper.GetString("ENV"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	e := echo.New()
	logger := lecho.New(
		os.Stdout,
		lecho.WithLevel(Log.DEBUG),
		lecho.WithTimestamp(),
		lecho.WithCaller(),
	)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Logger = logger
	e.Use(lecho.Middleware(lecho.Config{
		Logger: logger}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	timeoutContext := time.Duration(viper.GetInt("CONTEXT.TIMEOUT")) * time.Second

	serviceType := viper.GetString("ServiceType")
	if serviceType == "" {
		log.Error().Msg("Configuration Error: ServiceType not available")

	}
	var _deviceService model.IDeviceService
	var _registryService model.IRegistryService
	var client *mongo.Client
	var ctx context.Context
	var cancel context.CancelFunc
	if serviceType == "gcp" {
		gcpurl := viper.GetString("ENV_GCPPORT")
		if gcpurl == "" {
			log.Error().Msg("Configuration Error: ENV_GCPPORT address not available")

		}
		_deviceService = gcpService.NewDeviceService(gcpurl)
		_registryService = gcpService.NewRegistryService(gcpurl)

	} else if serviceType == "kore" {
		MongoCS := viper.GetString("MongoCS")
		if MongoCS == "" {
			log.Error().Msg("Configuration Error: MongoDB Connection String address not available")

		}
		MongoDB := viper.GetString("MongoDB")
		if MongoDB == "" {
			log.Error().Msg("Configuration Error: MongoDB Database String not available")

		}
		RegistryCollection := viper.GetString("RegistryCollection")
		if RegistryCollection == "" {
			log.Error().Msg("Configuration Error: MongoDB Registry Collection String not available")

		}
		DeviceCollection := viper.GetString("DeviceCollection")
		if DeviceCollection == "" {
			log.Error().Msg("Configuration Error: MongoDB Device Collection String not available")

		}
		var err error
		client, ctx, err = connect(MongoCS)
		if err != nil {
			panic(err)
		}
		ping(client, ctx)
		_deviceService = koreService.NewDeviceService(client, DeviceCollection, MongoDB, ctx)
		_registryService = koreService.NewRegistryService(client, RegistryCollection, MongoDB, ctx)

	} else {
		log.Fatal().Msg("Configuration Error: Service Type Not Found")
	}
	_deviceUsecase := iotUsecase.NewDeviceUsecase(_deviceService, timeoutContext)
	_registryUsecase := iotUsecase.NewIoTUsecase(_registryService, timeoutContext)
	defer func() {
		if serviceType == "kore" {
			close(client, ctx, cancel)
		}
	}()
	iotDelivery.NewIoTtHandler(e, _registryUsecase, _deviceUsecase)
	log.Fatal().Err(e.Start(viper.GetString("ENV_AUTH_SERVER"))).Msg("")

}
