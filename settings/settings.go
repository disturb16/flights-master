package settings

import (
	"encoding/json"
	"log"
	"os"
)

type SerpapiCredentials struct {
	ApiKey string `json:"api_key"`
}

type Settings struct {
	DbPath   string             `json:"db_path"`
	AppAddr  string             `json:"app_addr"`
	LogLevel int                `json:"log_level"`
	Serpapi  SerpapiCredentials `json:"serpapi"`
}

var defaultPath = "./conf.json"

func New() (*Settings, error) {
	bb, err := os.ReadFile(defaultPath)
	if err != nil {
		log.Println("failed to read")
		return nil, err
	}

	settings := &Settings{}
	err = json.Unmarshal(bb, settings)
	if err != nil {
		log.Println("failed to unmarshal")
		return nil, err
	}

	return settings, nil
}
