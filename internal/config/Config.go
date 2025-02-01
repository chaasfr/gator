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

func getConfigFilePath() (string, error) {
	folder,err := os.UserHomeDir()
	if err != nil {
		return "",fmt.Errorf("error finding home dir %w", err)
	}
	return fmt.Sprintf("%s/%s",folder,GatorConfigPath), nil
}

func Read() (*Config, error){
	url, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	configBytes, err := os.ReadFile(url)

	if err != nil {
		return nil, fmt.Errorf("cannot read config at %s: %w", url, err)
	}

	var conf Config
	if err := json.Unmarshal(configBytes, &conf); err != nil {
		return nil, fmt.Errorf("cannot parse config at %s: %w", url, err)
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

	url, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(url, confBytes, 0644)
	if err != nil {
		return fmt.Errorf("cannot save file %s - %w", url, err)
	}

	return nil
}