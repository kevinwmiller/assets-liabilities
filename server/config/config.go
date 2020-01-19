package config

import (
	"context"
	"time"

	"github.com/spf13/viper"
)

// DevelopmentMode indicates the backend is running in development mode
// This mode trades security for ease of use (Such as using unsecure cookies)
const DevelopmentMode = "development"

// ProductionMode indicates the backend is running in production mode
// This mode has the highest level of security
const ProductionMode = "production"

// DefaultMode is the default mode of the backend
const DefaultMode = ProductionMode

// Ctx is the context type used to reference the current configuration
type Ctx struct{}

// Configuration is the main configuration file loader. The format is expected to be a toml file
type Configuration struct {
	AppName         string
	Mode            string
	Domain          string
	AllowableOrigin string

	Address string
	Port    int32

	DBHost     string
	DBPort     int32
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string

	SessionSecret   string
	SessionExpireIn time.Duration

	DebugMode bool

	ReadTimeoutInSeconds  time.Duration
	WriteTimeoutInSeconds time.Duration
}

// LoadConfig loads the toml file located in the given path with the given basename
func LoadConfig(path string, configFileName string) *Configuration {
	config := viper.New()

	config.AutomaticEnv()
	config.SetEnvPrefix("AL")
	config.SetDefault("APP_NAME", "AssetsLiabilities")
	config.SetDefault("MODE", DefaultMode)
	config.SetDefault("DOMAIN", "https://kwmassetsliabilities.com")
	config.SetDefault("ALLOWABLE_ORIGIN", "kwmassetsliabilities.com")

	config.SetDefault("ADDRESS", "127.0.0.1")
	config.SetDefault("PORT", "8080")

	config.SetDefault("DB_HOST", "localhost")
	config.SetDefault("DB_PORT", "5432")
	config.SetDefault("DB_NAME", "postgres")
	config.SetDefault("DB_USER", "postgres")
	config.SetDefault("DB_PASSWORD", "postgres")
	config.SetDefault("DB_SSL_MODE", "disable")

	config.SetDefault("SESSION_SECRET", "")
	config.SetDefault("SESSION_EXPIRE_IN", "40000s")
	config.SetDefault("DEBUG_MODE", "false")

	config.SetDefault("READ_TIMEOUT_IN_SECONDS", "15s")
	config.SetDefault("WRITE_TIMEOUT_IN_SECONDS", "15s")

	return &Configuration{
		AppName:         config.GetString("APP_NAME"),
		Mode:            config.GetString("MODE"),
		Domain:          config.GetString("DOMAIN"),
		AllowableOrigin: config.GetString("ALLOWABLE_ORIGIN"),

		Address: config.GetString("ADDRESS"),
		Port:    config.GetInt32("PORT"),

		DBHost:     config.GetString("DB_HOST"),
		DBPort:     config.GetInt32("DB_PORT"),
		DBName:     config.GetString("DB_NAME"),
		DBUser:     config.GetString("DB_USER"),
		DBPassword: config.GetString("DB_PASSWORD"),
		DBSSLMode:  config.GetString("DB_SSL_MODE"),

		DebugMode: config.GetBool("DEBUG_MODE"),

		SessionSecret:   config.GetString("SESSION_SECRET"),
		SessionExpireIn: config.GetDuration("SESSION_EXPIRE_IN"),

		ReadTimeoutInSeconds:  config.GetDuration("READ_TIMEOUT_IN_SECONDS"),
		WriteTimeoutInSeconds: config.GetDuration("WRITE_TIMEOUT_IN_SECONDS"),
	}
}

// Config returns the context configuration
// A config object must be set in every context, so if one does not exist, panic
func Config(ctx context.Context) *Configuration {
	config, ok := ctx.Value(Ctx{}).(*Configuration)
	if !ok {
		panic("Error: config not set in context")
	}
	return config
}
