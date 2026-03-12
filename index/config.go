package index

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
)

const CONFIG_FILENAME string = "registry.config.json"

type Config struct {
	Input  ConfigInput `json:"input"`
	Output string      `json:"output"`
}

type ConfigInput struct {
	Folders []ConfigInputFolder `json:"folders"`
	Files   []ConfigInputFile   `json:"files"`
}

type ConfigInputFolder struct {
	Name    string                  `json:"name"`
	Path    string                  `json:"path"`
	Ignores ConfigInputFolderIgnore `json:"ignores"`
}

type ConfigInputFolderIgnore struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
}

type ConfigInputFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func GetIndexConfig() (Config, error) {
	config := Config{}

	cwd, err := os.Getwd()
	if err != nil {
		return config, err
	}

	configPath := path.Join(cwd, CONFIG_FILENAME)
	content, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	// TODO Validator
	err = json.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}

	if config.Output == "" {
		config.Output = ".registry"
	}

	if filepath.IsLocal(string(config.Output)) == false {
		return config, errors.New("config: Output path must be local")
	}

	outputAbsPath, err := filepath.Abs(string(config.Output))
	if err != nil {
		return config, err
	}

	config.Output = outputAbsPath

	return config, nil
}
