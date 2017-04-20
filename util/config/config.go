package config

import (
	"encoding/json"
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

type ConfigDB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Util Config :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func ReadConfig(cfg interface{}, path string, module string) error {
	environ := os.Getenv("MY-ENV")
	if environ == "" {
		environ = "development"
	}

	fname := path + "/" + module + "." + environ + ".ini"

	err := gcfg.ReadFileInto(cfg, fname)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfigJson(pathFile string) (ConfigDB, error) {
	var configDB ConfigDB
	configFile, err := os.Open(pathFile)
	if err != nil {
		return configDB, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configDB)
	if err != nil {
		return configDB, err
	}

	return configDB, nil

}

type Config struct {
	Database map[string]*database
}

type database struct {
	Master string
}
