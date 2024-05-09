package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TraceLogs bool   `envconfig:"TRACE_LOGS"`
	Port      string `envconfig:"PORT" required:"true"`
	// GcpCredsPath string `envconfig:"GCP_CREDS_PATH" required:"true"`
}

func NewConfig() *Config {
	config := &Config{
		TraceLogs: false,
	}
	return config
}

func (config *Config) InitConfig() {
	if err := envconfig.Process("myImageSvc", config); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
