package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config is an app's configuration
type Config struct {
	Port           int    `json:"port"`
	DBURL          string `json:"dbURL"`
	EncryptionKey  string `json:"encryptionKey"`
	Secret         string `json:"secret"`
	AuthURL        string `json:"authURL"`
	ConsumerSecret string `json:"consumerSecret"`
}

func load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	return cfg, err
}

// Dev return development config
func Dev() (*Config, error) {
	return load("config.dev.json")
}

// Prod return production config
func Prod() (*Config, error) {
	return load("config.prod.json")
}
