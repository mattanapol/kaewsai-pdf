package setting

import (
	"time"
)

type ApiConfiguration struct {
	Server   *Server  `mapstructure:"server"`
	Aws      *Aws     `mapstructure:"aws"`
	MongoDb  *MongoDb `mapstructure:"mongodb"`
	Wk       *SQS     `mapstructure:"wk"`
	Chromium *SQS     `mapstructure:"chromium"`
}

// NewApiConfiguration initialize the configuration instance
func NewApiConfiguration() (ApiConfiguration, error) {
	cfg := ApiConfiguration{
		Server:   &Server{},
		Aws:      &Aws{},
		MongoDb:  &MongoDb{},
		Wk:       &SQS{},
		Chromium: &SQS{},
	}

	readFile(cfg, "config", "api", "")

	cfg.Server.ReadTimeout = cfg.Server.ReadTimeout * time.Second
	cfg.Server.WriteTimeout = cfg.Server.WriteTimeout * time.Second

	return cfg, nil
}
