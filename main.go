package main

import (
	"os"

	"github.com/nenodias/go-config-client/config"
)

func main() {
	os.Setenv("SPRING_CLOUD_CONFIG", "http://localhost:8888")
	os.Setenv("SPRING_CLOUD_LABEL", "default")
	os.Setenv("SPRING_PROFILES_ACTIVE", "default")
	config.GetConfig("greendogdelivery")
}
