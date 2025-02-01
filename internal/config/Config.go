package config

import (
	"encoding/json"
	"fmt"
	"os"
)
const GatorConfigPath = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`           
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error){
	configBytes, err := os.ReadFile(GatorConfigPath)

	if err != nil {
		return nil, fmt.Errorf("cannot read config at %s: %w", GatorConfigPath, err)
	}

	var conf Config
	if err := json.Unmarshal(configBytes, &conf); err != nil {
		return nil, fmt.Errorf("cannot parse config at %s: %w", GatorConfigPath, err)
	}
	return &conf, nil
}

func (conf *Config) SetUser(username string) error {
	conf.CurrentUserName = username
	return write(conf)
}

func write(conf *Config) error {
	confBytes, err := json.Marshal(conf)
	if err != nil {
		return fmt.Errorf("cannot marshal config - %w", err)
	}

	err = os.WriteFile(GatorConfigPath, confBytes, 0644)
	if err != nil {
		return fmt.Errorf("cannot save file %s - %w", GatorConfigPath, err)
	}

	return nil
}