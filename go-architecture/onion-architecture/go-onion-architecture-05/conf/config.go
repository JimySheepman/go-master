package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
		Mock bool
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

var Current *Config

func NewConfig(runServer bool) {
	var C Config
	Current = &C
	viper.AddConfigPath("$GOPATH/go-echo-rest-sample/conf/")
	viper.SetConfigType("yml")

	if runServer {
		viper.SetConfigName("production")
	} else {
		viper.SetConfigName("local")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal config file error: %s", err))
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("fatal config file error: %s", err))
	}
	return
}
