package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func makeNewConfig(t *testing.T, filePath string) *Config {
	cfg, err := NewConfig(filePath)
	if err != nil {
		t.Fatalf("error: got=%v want=nil", err)
	}
	return cfg
}

func makeTempDir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		t.Fatalf("makeTempDir(): %v", err)
	}
	return dir, func() {
		err = os.RemoveAll(dir)
		if err != nil {
			t.Fatalf("makeTempDir(): %v", err)
		}
	}
}

func makeTempNewConfig(t *testing.T) (*Config, func()) {
	dir, close := makeTempDir(t)

	filePath := filepath.Join(dir, "config.toml")
	cfg := makeNewConfig(t, filePath)

	return cfg, close
}

func TestNewConfigWithDir(t *testing.T) {
	t.Parallel()

	// don't use `makeTempNewConfig()` here
	// because we want to control and compare filePath
	dir, close := makeTempDir(t)
	defer close()

	filePath := filepath.Join(dir, "config.toml")
	cfg := makeNewConfig(t, filePath)

	if cfg.dirPath != dir {
		t.Errorf("cfg.fileName: got=%v want=%v", cfg.dirPath, dir)
	}
	if cfg.fileName != "config.toml" {
		t.Errorf("cfg.fileName: got=%v want=%v", cfg.fileName, "config.toml")
	}
	if cfg.filePath != filePath {
		t.Errorf("cfg.fileName: got=%v want=%v", cfg.filePath, filePath)
	}
}

func TestReadEmptyDirectory(t *testing.T) {
	t.Parallel()

	cfg, close := makeTempNewConfig(t)
	defer close()

	cfg.Read()

	if cfg.GithubToken != "" {
		t.Errorf("cfg.GithubToken: got=%v want=%v", cfg.GithubToken, "")
	}
}

func TestReadMissingDirectory(t *testing.T) {
	t.Parallel()

	// don't use `makeTempNewConfig()` here
	// because we want to control and compare filePath
	dir, close := makeTempDir(t)
	defer close()

	filePath := filepath.Join(dir, "missing", "config.toml")
	cfg := makeNewConfig(t, filePath)

	cfg.Read()

	if cfg.GithubToken != "" {
		t.Errorf("cfg.GithubToken: got=%v want=%v", cfg.GithubToken, "")
	}
}

func TestWriteReadEmptyDirectory(t *testing.T) {
	t.Parallel()

	// don't use `makeTempNewConfig()` here
	// because we want to reuse filePath
	dir, close := makeTempDir(t)
	defer close()

	filePath := filepath.Join(dir, "config.toml")
	cfg := makeNewConfig(t, filePath)

	cfg.GithubToken = "abcd1234"
	err := cfg.Write()
	if err != nil {
		t.Errorf("cfg.Write(): %v", err)
	}

	cfg = makeNewConfig(t, filePath)
	cfg.Read()

	if cfg.GithubToken != "abcd1234" {
		t.Errorf("cfg.GithubToken: got=%v want=%v", cfg.GithubToken, "")
	}
}

func TestWriteReadMissingDirectory(t *testing.T) {
	t.Parallel()

	// don't use `makeTempNewConfig()` here
	// because we want to control and reuse filePath
	dir, close := makeTempDir(t)
	defer close()

	filePath := filepath.Join(dir, "missing", "config.toml")
	cfg := makeNewConfig(t, filePath)

	cfg.GithubToken = "abcd1234"
	err := cfg.Write()
	if err != nil {
		t.Errorf("cfg.Write(): %v", err)
	}

	cfg = makeNewConfig(t, filePath)
	cfg.Read()

	if cfg.GithubToken != "abcd1234" {
		t.Errorf("cfg.GithubToken: got=%v want=%v", cfg.GithubToken, "")
	}
}
