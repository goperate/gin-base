package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

var CONF *viper.Viper

func init() {
	CONF = viper.New()
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	if os.Getenv("APP_MODE") == "prod" {
		CONF.SetConfigName("prod")
	} else {
		CONF.SetConfigName("dev")
	}
	CONF.SetConfigType("yaml")
	fmt.Println(path)
	CONF.AddConfigPath(path)
	err := CONF.ReadInConfig()
	if err != nil {
		panic("配置文件加载失败! " + err.Error())
	}
	ConnectDB()
}
