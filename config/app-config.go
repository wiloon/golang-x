package config

import (
	"wiloon.com/golang-x/config/system"
	"log"
	"github.com/go-akka/configuration"
)

func GetString(key string) string {
	sysEnvXConfig := system.GetSystemEnv("X_CONFIG")
	log.Println("sys env x config:", sysEnvXConfig)
	conf := configuration.LoadConfig(sysEnvXConfig)
	value := conf.GetString(key);
	log.Printf("key:%s, value:%s", key, value)
	return value

}
