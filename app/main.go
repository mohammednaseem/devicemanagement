package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"

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
	var _deviceUsecase model.IDevicerUsecase
	var _registryUsecase model.IRegistryrUsecase
	if serviceType == "gcp" {
		gcpurl := viper.GetString("ENV_GCPPORT")
		if gcpurl == "" {
			log.Error().Msg("Configuration Error: ENV_GCPPORT address not available")

		}
		_deviceService = gcpService.NewDeviceService(gcpurl)
		_registryService = gcpService.NewRegistryService(gcpurl)
		_deviceUsecase = iotUsecase.NewDeviceUsecase(_deviceService, timeoutContext)
		_registryUsecase = iotUsecase.NewIoTUsecase(_registryService, timeoutContext)

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
	} else {
		log.Fatal().Msg("Configuration Error: Service Type Not Found")
	}

	iotDelivery.NewIoTtHandler(e, _registryUsecase, _deviceUsecase)

	log.Fatal().Err(e.Start(viper.GetString("ENV_AUTH_SERVER"))).Msg("")
}
