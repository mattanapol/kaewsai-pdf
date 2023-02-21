package setting

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type ApiConfiguration struct {
	Server *Server `mapstructure:"server"`
}

// NewApiConfiguration initialize the configuration instance
func NewApiConfiguration() (ApiConfiguration, error) {
	cfg := ApiConfiguration{
		Server: &Server{},
	}

	readFile(cfg)

	cfg.Server.ReadTimeout = cfg.Server.ReadTimeout * time.Second
	cfg.Server.WriteTimeout = cfg.Server.WriteTimeout * time.Second

	return cfg, nil
}

func readFile(cfg interface{}) {
	ymlConfig := viper.New()
	ymlConfig.AddConfigPath("config")
	ymlConfig.SetConfigName("api")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	ymlConfig.SetEnvKeyReplacer(replacer)
	ymlConfig.SetEnvPrefix("API")
	ymlConfig.AutomaticEnv()

	err := ymlConfig.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	err = ymlConfig.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}
