package system

import (
	"testing"
	"fmt"
)

func TestGetSystemEnv(t *testing.T) {
	fmt.Println(GetSystemEnv("JAVA_HOME"))
	fmt.Println(GetSystemEnv("X_CONFIG"))
}
