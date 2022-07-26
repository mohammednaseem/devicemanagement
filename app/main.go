package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"

	iotDelivery "github.com/gcp-iot/implementation/delivery/http"
	iotService "github.com/gcp-iot/implementation/service/gcp"
	iotUsecase "github.com/gcp-iot/implementation/usecase"

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

	gcpurl := viper.GetString("ENV_GCPPORT")
	if gcpurl == "" {
		gcpurl = viper.GetString(`gcp_port`)
	}

	if gcpurl == "" {
		log.Error().Msg("Configuration Error: ENV_PPSA address not available")
	}

	_deviceService := iotService.NewDeviceService(gcpurl)
	_deviceUsecase := iotUsecase.NewDeviceUsecase(_deviceService, timeoutContext)
	_registryService := iotService.NewRegistryService(gcpurl)
	_registryUsecase := iotUsecase.NewIoTUsecase(_registryService, timeoutContext)
	iotDelivery.NewIoTtHandler(e, _registryUsecase, _deviceUsecase)

	log.Fatal().Err(e.Start(viper.GetString("ENV_AUTH_SERVER"))).Msg("")
}
