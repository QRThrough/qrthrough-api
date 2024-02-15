package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config LocalConfig

func InitConfig() {
	mode := os.Getenv("mode")
	if mode == "" {
		mode = "dev"
	}
	env := fmt.Sprintf("./config/.env.%v",mode)
	
	viper.SetConfigFile(env)
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	readConfig()
}

func readConfig() {
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
