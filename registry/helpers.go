package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"glitchkids/registry_gk/helpers"
	"glitchkids/registry_gk/index"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
)

func getConfigPath() (string, error) {
	var configPath string

	configDir, err := os.UserConfigDir()
	if err != nil {
		return configPath, err
	}

	configPath = filepath.Join(configDir, config_path)

	return configPath, nil
}

func writeConfig(remoteRegistry RemoteRegistry) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	dirname := filepath.Dir(path)
	if helpers.Access(dirname) == false {
		if err := helpers.Mkdir(dirname); err != nil {
			return err
		}
	}

	content, err := helpers.ToJSON(remoteRegistry)
	if err != nil {
		return err
	}

	if err := helpers.WriteFile(path, []byte(content)); err != nil {
		return err
	}

	return nil
}

func readConfig() (RemoteRegistry, error) {
	config := RemoteRegistry{}

	path, err := getConfigPath()
	if err != nil {
		return config, err
	}

	if helpers.Access(path) == false {
		return config, errors.New("Can't Access Config")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func getRegistryItemIndexByName(remoteRegistry RemoteRegistry, name string) (int, error) {
	index := slices.IndexFunc(remoteRegistry, func(e RemoteRegistryItem) bool {
		return e.Name == name
	})

	if index == -1 {
		return index, fmt.Errorf("Registry not found : %v", name)
	}

	return index, nil
}

func getRegistryItemByName(remoteRegistry RemoteRegistry, name string) (RemoteRegistryItem, error) {
	index, err := getRegistryItemIndexByName(remoteRegistry, name)
	if err != nil {
		return RemoteRegistryItem{}, err
	}

	return remoteRegistry[index], nil
}

func fetchJSON[T index.RegistryIndex | index.RegistryFolder](url string, ref *T) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, ref)
	if err != nil {
		return err
	}

	return nil
}

func fetchRegistryIndex(baseUrl string, ref *index.RegistryIndex) error {
	finalUrl := fmt.Sprintf("%v/index.json", baseUrl)
	return fetchJSON(finalUrl, ref)
}

func fetchRegistryItem[T index.RegistryFolder](baseUrl string, filename string, ref *T) error {
	finalUrl := fmt.Sprintf("%v/%v", baseUrl, filename)
	return fetchJSON(finalUrl, ref)
}

func writeRegistryFiles(files []index.RegistryFile, outDir string, force bool) error {
	if helpers.Access(outDir) == false {
		if err := helpers.Mkdir(outDir); err != nil {
			return err
		}
	}

	for _, v := range files {
		path := filepath.Join(outDir, v.Path)

		if helpers.Access(filepath.Dir(path)) == true && force == false {
			// TODO ignore or ovewrite prompt
		}

		dir := filepath.Dir(path)
		if helpers.Access(dir) == false {
			if err := helpers.Mkdir(dir); err != nil {
				return err
			}
		}

		helpers.WriteFile(path, []byte(v.Content))
	}

	return nil
}
