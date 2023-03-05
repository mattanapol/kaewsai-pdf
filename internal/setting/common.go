package setting

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func readFile(cfg interface{}, configPath string, configName string, envPrefix string) {
	ymlConfig := viper.New()
	ymlConfig.AddConfigPath(configPath)
	ymlConfig.SetConfigName(configName)
	replacer := strings.NewReplacer(".", "_", "-", "_")
	ymlConfig.SetEnvKeyReplacer(replacer)
	if envPrefix != "" {
		ymlConfig.SetEnvPrefix(envPrefix)
	}
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
