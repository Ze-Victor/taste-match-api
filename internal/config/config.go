package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

// Configuration is a global variable that stores all mapped configuration
var configuration *Config

// Config is a struct that represents all Configuration structure
type Config struct {
	AppName  string `mapstructure:"APP_NAME" default:"${{ taste-match-apirtifactId }}"`
	Version  string `mapstructure:"APP_VERSION" default:"1.0.0"`
	Env      string `mapstructure:"GO_ENV"`
	Api      api
	Server   server
	Database database
}

type api struct {
}

type server struct {
	Host                   string `mapstructure:"SERVER_HOST"`
	Port                   string `mapstructure:"SERVER_PORT"`
	TimeToCompleteShutdown int    `mapstructure:"SERVER_TIME_TO_COMPLETE_SHUTDOWN" default:"60"`
}

type database struct {
	User          string `mapstructure:"DB_USER"`
	Pass          string `mapstructure:"DB_PASS"`
	UserMigration string `mapstructure:"DB_USER_MIGRATION"`
	PassMigration string `mapstructure:"DB_PASS_MIGRATION"`
	Host          string `mapstructure:"DB_HOST"`
	Port          string `mapstructure:"DB_PORT"`
	Name          string `mapstructure:"DB_NAME"`
	EnableLog     bool   `mapstructure:"DB_ENABLE_LOG"`
}

func getMappedEnvs(configStruct reflect.Type) []string {
	result := make([]string, 0)

	for i := 0; i < configStruct.NumField(); i++ {
		field := configStruct.Field(i)
		if configName := field.Tag.Get("mapstructure"); configName != "" {
			result = append(result, configName)
		}
		if field.Type.Kind() == reflect.Struct {
			result = append(result, getMappedEnvs(field.Type)...)
		}
	}
	return result
}

func setDefaultValues(configStruct reflect.Type) {
	for i := 0; i < configStruct.NumField(); i++ {
		field := configStruct.Field(i)
		configName := field.Tag.Get("mapstructure")
		defaultValue := field.Tag.Get("default")

		if configName != "" && defaultValue != "" {
			viper.SetDefault(configName, defaultValue)
		}

		if field.Type.Kind() == reflect.Struct {
			setDefaultValues(field.Type)
		}
	}
}

// Load is the method responsible for fill Config Structure
func Load() error {
	configuration = &Config{}

	environment := os.Getenv("GO_ENV")
	if environment == "" {
		fmt.Println("[Method: Config.Load()] Your GO_ENV was not filled, configure it on environment or env file and try again.")
	}
	envFile := ".env-"
	envPath := "."
	if environment == "" {
		envFile += "development"
	} else if environment == "test" {
		envFile = envFile + environment
		envPath = "../../test/"
	} else {
		envFile += environment
	}

	viper.AddConfigPath(envPath)
	viper.SetConfigName(envFile)
	viper.SetConfigType("env")

	setDefaultValues(reflect.TypeOf(Config{}))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("[Method: Config.Load()]", envFile, "not found, load by environment variables")

			viper.AutomaticEnv()
			mapped := getMappedEnvs(reflect.TypeOf(Config{}))
			for _, env := range mapped {
				viper.BindEnv(env)
			}

		} else {
			return err
		}
	} else {
		fmt.Println("[Method: Config.Load()] Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return err
	}

	if err := viper.Unmarshal(&configuration.Api); err != nil {
		return err
	}

	if err := viper.Unmarshal(&configuration.Server); err != nil {
		return err
	}

	if err := viper.Unmarshal(&configuration.Database); err != nil {
		return err
	}

	return nil

}

// Get returns a Config Structure
func Get() *Config {
	return configuration
}
