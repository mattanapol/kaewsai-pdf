package setting

import "time"

type Server struct {
	RunMode      string        `mapstructure:"run-mode"`
	HttpPort     int           `mapstructure:"http-port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
}
