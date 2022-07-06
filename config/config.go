package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PropertySource struct {
	Name   string         `json:"name"`
	Source map[string]any `json:"source"`
}

type ConfigProperties struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           string           `json:"label"`
	Version         string           `json:"version"`
	State           string           `json:"state"`
	PropertySources []PropertySource `json:"propertySources"`
}

var Config = ConfigProperties{}

func GetConfig(app string) error {
	cloudURI := os.Getenv("SPRING_CLOUD_URI")
	label := os.Getenv("SPRING_CLOUD_LABEL")
	profiles := os.Getenv("SPRING_PROFILES_ACTIVE")
	url := fmt.Sprintf("%s/%s/%s/%s", cloudURI, app, profiles, label)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	req.Header.Add("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	buf := &bytes.Buffer{}
	tee := io.TeeReader(res.Body, buf)
	jsonDecoder := json.NewDecoder(tee)
	err = jsonDecoder.Decode(&Config)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	dados, err := ioutil.ReadAll(buf)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	log.Print(string(dados))
	return nil
}
