package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetConfig(app string) error {
	cloudURI := os.Getenv("SPRING_CLOUD_URI")
	label := os.Getenv("SPRING_CLOUD_LABEL")
	profiles := os.Getenv("SPRING_PROFILES_ACTIVE")
	url := fmt.Sprintf("%s/%s/%s/%s", cloudURI, app, profiles, label)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	dados, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	log.Print(string(dados))
	return nil
}
