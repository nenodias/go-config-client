package main

import (
	"fmt"
	"os"

	"github.com/nenodias/go-config-client/config"
)

type Prefixable interface {
	Prefix() string
}

type EurekaInstance struct {
	NonSecurePort   int64 `config:"nonSecurePort"`
	PreferIpAddress bool  `config:"preferIpAddress"`
}

func (e *EurekaInstance) Prefix() string {
	return "instance"
}

type EurekaProperties struct {
	Instance    EurekaInstance `config:"instance"`
	DefaultZone string         `config:"client.serviceUrl.defaultZone"`
}

func (e *EurekaProperties) Prefix() string {
	return "eureka"
}

func main() {
	os.Setenv("SPRING_CLOUD_URI", "http://localhost:8888")
	os.Setenv("SPRING_CLOUD_LABEL", "main")
	os.Setenv("SPRING_PROFILES_ACTIVE", "default,dev")
	config.GetConfig("test")

	props := make(map[string]any)
	for _, v := range config.Config.PropertySources {
		for k := range v.Source {
			if _, ok := props[k]; !ok {
				props[k] = v.Source[k]
			}
		}
	}

	q := EurekaProperties{}
	fmt.Println(q)

}
