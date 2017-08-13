package config

import (
	"bufio"
	"errors"
	"fmt"
	"log"
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

	IsDirty bool

	dirPath  string
	fileName string
	filePath string
}

/*
NewConfig initialises a Config for you,
using defaults if you do not supply `filePath`
*/
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
	if err != nil {
		// in case of error, use defaults
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("error closing file after read %v: %v", c.filePath, err)
		}
	}()

	fileReader := bufio.NewReader(file)
	_, err = toml.DecodeReader(fileReader, c)
	if err != nil {
		// in case of error, use defaults
		return
	}
}

func (c *Config) Write() error {
	err := os.MkdirAll(c.dirPath, os.FileMode(0700))
	if err != nil {
		return fmt.Errorf("error creating directories: %v %v", c.dirPath, err)
	}

	file, err := os.OpenFile(c.filePath, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		return fmt.Errorf("error creating file: %v %v", c.filePath, err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("error closing file after write %v: %v", c.filePath, err)
		}
	}()

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
