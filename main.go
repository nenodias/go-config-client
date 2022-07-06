package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/nenodias/go-config-client/config"
	"github.com/nenodias/go-config-client/load"
)

type Datasource struct {
	PoolName    string `config:"pool-name"`
	IdleTimeout int64  `config:"idle-timeout"`
}

func (e Datasource) Prefix() string {
	return "hikari"
}

type SpringDatasource struct {
	Datasource Datasource `config:"datasource"`
	ShowSQL    bool       `config:"jpa.show-sql"`
}

func (e SpringDatasource) Prefix() string {
	return "spring"
}

func main() {
	os.Setenv("SPRING_CLOUD_URI", "https://localhost:8888")
	os.Setenv("SPRING_CLOUD_LABEL", "qa")
	os.Setenv("SPRING_PROFILES_ACTIVE", "qa")
	user := ""
	pass := ``
	auth := ""
	if user != "" && pass != "" {
		auth = base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	}

	config.GetConfig("test", auth)

	props := make(map[string]any)
	for _, v := range config.Config.PropertySources {
		for k := range v.Source {
			if _, ok := props[k]; !ok {
				props[k] = v.Source[k]
			}
		}
	}
	q := SpringDatasource{}
	load.Properties(&q, props)
	fmt.Println(q.ShowSQL)
	fmt.Println(q.Datasource.IdleTimeout)
	fmt.Println(q.Datasource.PoolName)
}
