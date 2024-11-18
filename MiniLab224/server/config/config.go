package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
)

type Config struct {
	ListenPort       string `json:"SERVER_LISTEN_PORT"`
	DatabaseUser     string `json:"DATABASE_USER"`
	DatabaseName     string `json:"DATABASE_NAME"`
	DatabasePassword string `json:"DATABASE_PASSWORD"`
	DatabaseHost     string `json:"DATABASE_HOST"`
	DatabasePort     string `json:"DATABASE_PORT"`
}

func (c *Config) GetPostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName,
	)
}

func newConfig() *Config {
	var myEnv map[string]string
	myEnv, err := godotenv.Read("../.env")

	if err != nil {
		return nil
	}

	var settings Config
	result, err := json.Marshal(myEnv)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(result, &settings)
	if err != nil {
		return nil
	}
	return &settings
}

var Settings = newConfig()
