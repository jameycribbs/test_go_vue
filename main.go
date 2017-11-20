package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port       string `json:"port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

func main() {
	var cfg Config

	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Cannot read config file!!!")
		return
	}

	err = json.Unmarshal(configData, &cfg)
	if err != nil {
		fmt.Println("Invalid config file!!!")
		return
	}

	a := App{}
	a.Initialize(cfg.DBUser, cfg.DBPassword, cfg.DBName)
	a.Run(":" + cfg.Port)
}
