package config

import (
	"wiloon.com/golang-x/config/system"
	"log"
	"github.com/go-akka/configuration"
	"path/filepath"
)

const SYS_ENV_KEY_APP_CONFIG = "app_config"
const DEFAULT_FILE_NAME = "app.config"

func GetString(key string) string {
	appConfigPath := system.GetSystemEnv(SYS_ENV_KEY_APP_CONFIG)
	if appConfigPath == "" {
		log.Fatalf("sys env '%s' not found", SYS_ENV_KEY_APP_CONFIG)
		return ""
	}

	fullPath := filepath.Join(appConfigPath, DEFAULT_FILE_NAME)
	log.Println("app config full path:", fullPath)

	conf := configuration.LoadConfig(fullPath)
	value := conf.GetString(key);
	log.Printf("key:%s, value:%s", key, value)
	return value

}
