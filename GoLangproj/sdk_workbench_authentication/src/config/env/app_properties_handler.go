package env

/*
	The purpose of this class is to consolidate all the application properties from either
	environment variables or from env file
*/

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Properties struct {
	AppName                      string `mapstructure:"APP_NAME"`
	AppDesc                      string `mapstructure:"APP_DESC"`
	AppCopyright                 string `mapstructure:"APP_COPYRIGHT"`
	AppEnv                       string `mapstructure:"APP_ENV"`
	LoggingLevel                 string `mapstructure:"LOGGING_LEVEL"`
	ResponseTimeout              int    `mapstructure:"RESPONSE_TIMEOUT"`
	BindIp                       string `mapstructure:"BIND_IP"`
	Port                         string `mapstructure:"PORT"`
	GinMode                      string `mapstructure:"GIN_MODE"`
	ArangoAddr                   string `mapstructure:"ARANGO_ADDR"`
	ArangoUser                   string `mapstructure:"ARANGO_USER"`
	ArangoPwd                    string `mapstructure:"ARANGO_PASSWORD"`
	ArangoDbName                 string `mapstructure:"ARANGO_DB_NAME"`
	ArangoCollectionName         string `mapstructure:"ARANGO_DB_COLLECTION_NAME"`
	ArangoHistoryCollectionName  string `mapstructure:"ARANGO_DB_HISTORY_COLLECTION"`
	ArangoDbCreateMode           string `mapstructure:"ARANGO_DB_CREATION_MODE"`
	MaxProcessInstancesPerEngine int    `mapstructure:"MAX_PROCESS_INSTANCES_PER_ENGINE"`
	MaxEngineInstances           int    `mapstructure:"MAX_ENGINE_INSTANCE"`
	EnableZBM                    bool   `mapstructure:"ENABLE_ZBM"`
	HazelcastHost                string `mapstructure:"HAZELCAST_HOST"`
	HazelcastPort                string `mapstructure:"HAZELCAST_PORT"`
}

var props Properties

func NewProperties() *Properties {

	log.SetOutput(os.Stderr)

	props = Properties{}

	viper.SetConfigFile(".env")

	// automatically override file variables if environment variables are present
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&props)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if props.AppEnv == "DEV" {
		log.Println("The App is running in development env")
	}

	return &props
}

// GetProperties() can be used globally to get application propetries
func GetProperties() Properties {
	return props
}
