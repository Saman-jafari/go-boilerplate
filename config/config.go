package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("../")
	config.AutomaticEnv()
	err = config.ReadInConfig()
	mapEnvToConfig(config)
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}

func mapEnvToConfig(config *viper.Viper) {
	allConfigs := config.AllKeys()
	for _, s := range allConfigs {
		envConf := strings.ToUpper(strings.Replace(s, ".", "_", -1))
		err := config.BindEnv(s, envConf)
		if err != nil {
			return
		}
	}

}
