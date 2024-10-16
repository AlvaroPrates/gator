package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	filename = ".gatorconfig.json"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username
	return write(*c)
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("when retrieving config filepath: %v", err)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}

	decoder := json.NewDecoder(file)
	var cfg Config

	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, filename), nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
