package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	configFilename = "config.toml"
)

// Config is how we persist / restore settings
type Config struct {
	GithubToken string `toml:"github_token"`
}

func (c *Config) Read() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configFilename)

	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		// it's okay, return the zero value of Config
		return nil
	}

	fileReader := bufio.NewReader(file)
	_, err = toml.DecodeReader(fileReader, c)
	if err != nil {
		// it's okay, return the zero value of Config
		return nil
	}

	return nil
}

func (c *Config) Write() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configFilename)

	err = os.MkdirAll(configDir, os.FileMode(0700))
	if err != nil {
		return fmt.Errorf("error creating directories: %v %v", configDir, err)
	}

	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
	defer file.Close()
	if err != nil {
		return fmt.Errorf("error creating file: %v %v", configPath, err)
	}

	fileWriter := bufio.NewWriter(file)
	encoder := toml.NewEncoder(fileWriter)
	err = encoder.Encode(c)
	if err != nil {
		return fmt.Errorf("error writing config: %v %v", configPath, err)
	}

	return nil
}

func getConfigDir() (string, error) {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		return "", errors.New("no HOME environment variable found")
	}
	return filepath.Join(home, ".config", "jokeyrhyme", "my-github-repos"), nil
}
