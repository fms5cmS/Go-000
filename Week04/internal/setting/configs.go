package setting

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host     string
	Name     string
	User     string
	Password string
}

func init() {
	viper.AddConfigPath("./Week04/configs/")
	viper.SetConfigName("app")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Can't find config file")
		} else {
			log.Fatal("Fatal error: ", err.Error())
		}
	}
	if err := viper.UnmarshalKey("database", &dataConfig); err != nil {
		log.Fatal("Parse database config error: ", err.Error())
	}
}
