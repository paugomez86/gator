package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/paugomez86/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *Config
}

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName string = ".gatorconfig.json"

func Read() (Config, error) {
	var config Config

	filePath, err := getConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("Unable to get config file path: %v", err)
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("Unable to read config file: %v", err)
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return config, fmt.Errorf("Unable to decode config file: %v", err)
	}

	return config, nil
}

func (c Config) SetUser(username string) error {
	c.CurrentUserName = username

	jsonData, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Unable to enconde config data: %v", err)
	}

	if err := write(jsonData); err != nil {
		return err
	}

	return nil
}

func write(data []byte) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Unable to find config file: %v", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("Unable to write to config file: %v", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	var filePath string

	if homeDir, err := os.UserHomeDir(); err != nil {
		return filePath, fmt.Errorf("Unable to find user home: %v", err)
	} else {
		filePath = filepath.Join(homeDir, configFileName)
	}

	return filePath, nil
}
