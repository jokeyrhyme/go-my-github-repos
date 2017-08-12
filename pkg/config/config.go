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
	defaultFileName = "config.toml"
)

// Config is how we persist / restore settings
type Config struct {
	GithubToken string `toml:"github_token"`

	dirPath  string
	fileName string
	filePath string
}

func NewConfig(filePath string) (*Config, error) {
	if filePath == "" {
		configDir, err := getDefaultConfigDir()
		if err != nil {
			return nil, err
		}

		return &Config{
			dirPath:  configDir,
			fileName: defaultFileName,
			filePath: filepath.Join(configDir, defaultFileName),
		}, nil
	}

	return &Config{
		dirPath:  filepath.Dir(filePath),
		fileName: filepath.Base(filePath),
		filePath: filePath,
	}, nil
}

func (c *Config) Read() {
	file, err := os.Open(c.filePath)
	defer file.Close()
	if err != nil {
		// it's okay, return the zero value of Config
	}

	fileReader := bufio.NewReader(file)
	_, err = toml.DecodeReader(fileReader, c)
	if err != nil {
		// it's okay, return the zero value of Config
	}
}

func (c *Config) Write() error {
	err := os.MkdirAll(c.dirPath, os.FileMode(0700))
	if err != nil {
		return fmt.Errorf("error creating directories: %v %v", c.dirPath, err)
	}

	file, err := os.OpenFile(c.filePath, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
	defer file.Close()
	if err != nil {
		return fmt.Errorf("error creating file: %v %v", c.filePath, err)
	}

	fileWriter := bufio.NewWriter(file)
	encoder := toml.NewEncoder(fileWriter)
	err = encoder.Encode(c)
	if err != nil {
		return fmt.Errorf("error writing config: %v %v", c.filePath, err)
	}

	return nil
}

func getDefaultConfigDir() (string, error) {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		return "", errors.New("no HOME environment variable found")
	}
	return filepath.Join(home, ".config", "jokeyrhyme", "my-github-repos"), nil
}
