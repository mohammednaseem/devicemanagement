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

	iotDelivery "github.com/RacoWireless/iot-gw-thing-management/implementation/_start/http"
	//gcpService "github.com/RacoWireless/iot-gw-thing-management/implementation/service/gcp"
	koreService "github.com/RacoWireless/iot-gw-thing-management/implementation/service/kore"
	iotUsecase "github.com/RacoWireless/iot-gw-thing-management/implementation/usecase"
	"github.com/RacoWireless/iot-gw-thing-management/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/spf13/viper"

	Log "github.com/labstack/gommon/log"
	"github.com/rs/zerolog/log"
	lecho "github.com/ziflex/lecho"
)

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

// @title IOT Device Management API
// @version 1.0
// @description This is a Iot Device Management  server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.korewireless.com
// @contact.email support@korewireless.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host iot.korewireless.com
// @BasePath /
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

	serviceType := viper.GetString("ENV_SERVICE_TYPE")
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
		//_deviceService = gcpService.NewDeviceService(gcpurl)
		//_registryService = gcpService.NewRegistryService(gcpurl)

	} else if serviceType == "kore" {
		MongoCS := viper.GetString("ENV_MONGO_CS")
		if MongoCS == "" {
			log.Error().Msg("Configuration Error: MongoDB Connection String address not available")

		}
		MongoDB := viper.GetString("ENV_MONGO_DB")
		if MongoDB == "" {
			log.Error().Msg("Configuration Error: MongoDB Database String not available")

		}
		RegistryCollection := viper.GetString("ENV_REGISTRY_COLLECTION")
		if RegistryCollection == "" {
			log.Error().Msg("Configuration Error: MongoDB Registry Collection String not available")

		}
		DeviceCollection := viper.GetString("ENV_DEVICE_COLLECTION")
		if DeviceCollection == "" {
			log.Error().Msg("Configuration Error: MongoDB Device Collection String not available")

		}
		Publish := viper.GetBool("ENV_PUBLISH")
		PubTopic := viper.GetString("ENV_PUB_TOPIC")
		if PubTopic == "" && Publish {
			log.Error().Msg("Configuration Error: PubTopic not available")

		}
		var err error
		ctx, client, err = koreService.Connect(MongoCS)
		if err != nil {
			panic(err)
		}
		koreService.Ping(ctx, client)
		_deviceService = koreService.NewDeviceService(ctx, client, DeviceCollection, RegistryCollection, MongoDB, PubTopic, Publish)
		_registryService = koreService.NewRegistryService(ctx, client, RegistryCollection, MongoDB, PubTopic, Publish)

	} else {
		log.Fatal().Msg("Configuration Error: Service Type Not Found")
	}
	_deviceUsecase := iotUsecase.NewDeviceUsecase(_deviceService, timeoutContext)
	_registryUsecase := iotUsecase.NewIoTUsecase(_registryService, timeoutContext)
	defer func() {
		if serviceType == "kore" {
			koreService.CloseMongo(ctx, client, cancel)
		}
	}()
	iotDelivery.NewIoTtHandler(e, _registryUsecase, _deviceUsecase)
	log.Error().Err(e.Start(viper.GetString("ENV_AUTH_SERVER"))).Msg("")

}
