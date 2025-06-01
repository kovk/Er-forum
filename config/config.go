package config

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	Server struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Dsn      string
	}
	JwtKey string
	App    struct {
		Host string
		Port string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}
}

var AppConfig *config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config initionalizing error:%v", err)
	}
	AppConfig = &config{}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("config initionalizing error:%v", err)
	}
	AppConfig.Database.Dsn = AppConfig.Database.User + ":" + AppConfig.Database.Password + "@tcp(" + AppConfig.Database.Host + AppConfig.Database.Port + ")/" + AppConfig.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	InitDB()
	InitRedis()
}
