package system

import "os"

func GetSystemEnv(key string) (value string) {
	return os.Getenv(key)
}
