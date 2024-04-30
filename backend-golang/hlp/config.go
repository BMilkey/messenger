package hlp

import (
	"os"
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var config AppConfig

type AppConfig struct {
	Title string `yaml:"title" envconfig:"TITLE" required:"true"`
	Http  HttpConfig
	Db    DatabaseConfig
}

type HttpConfig struct {
	Port string `yaml:"port" envconfig:"HTTP_PORT" required:"true"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" envconfig:"DB_HOST" required:"true"`
	Port     string `yaml:"port" envconfig:"DB_PORT" required:"true"`
	User     string `yaml:"user" envconfig:"DB_USER" required:"true"`
	Password string `yaml:"password" envconfig:"DB_PASSWORD" required:"true"`
	DbName   string `yaml:"dbname" envconfig:"DB_NAME" required:"true"`
}

func GetConfig(filepath string) (AppConfig, error) {
	if config == (AppConfig{}) {
		// Load default config file
		config.readYaml("config.yaml")
		// Load specific configuration with env variables
		config.readEnv()
	}

	configValue := reflect.ValueOf(config)
	return config, CheckStructEmptyFields(configValue, "")
}

func (config *AppConfig) readEnv() {
	envconfig.Process("app", config)
}
func (config *AppConfig) readYaml(filepath string) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error Reading Config file with path: %v\n", filepath)
	}
	yaml.UnmarshalStrict(yamlFile, config)
}
