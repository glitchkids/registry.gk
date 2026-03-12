package registry

import (
	"errors"
	"fmt"
	"glitchkids/registry_gk/helpers"
	"glitchkids/registry_gk/index"
	"path/filepath"
	"slices"
)

type RemoteRegistryItem struct {
	Url  string
	Name string
}
type RemoteRegistry []RemoteRegistryItem

const config_path string = "glitchkids/remote_registry.json"

func (r RemoteRegistryItem) String() string {
	return fmt.Sprintf("%v => %v\n", r.Name, r.Url)
}

func (r RemoteRegistry) String() string {
	var s string

	if len(r) == 0 {
		return "No remote registry added."
	}

	for i, v := range r {
		s = fmt.Sprintf("%v%v. %v\n", s, i, v)
	}
	return s
}

func GetOrCreateRemoteRegistry() (RemoteRegistry, error) {
	config := RemoteRegistry{}

	configPath, err := getConfigPath()
	if err != nil {
		return config, err
	}

	if helpers.Access(configPath) == false {
		if err = writeConfig(config); err != nil {
			return config, err
		}
	}

	return readConfig()
}

func AddItemToRemoteRegistry(remoteRegistry RemoteRegistry, name string, url string) (RemoteRegistryItem, error) {
	remoteRegistryItem := RemoteRegistryItem{}

	if _, err := getRegistryItemIndexByName(remoteRegistry, name); err == nil {
		return remoteRegistryItem, errors.New("Registry already exists")
	}

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	remoteRegistryItem.Name = name
	remoteRegistryItem.Url = url

	remoteRegistry = append(remoteRegistry, remoteRegistryItem)

	err := writeConfig(remoteRegistry)
	if err != nil {
		return remoteRegistryItem, err
	}

	return remoteRegistryItem, nil
}

func RemoveItemToRemoteRegistryByName(remoteRegistry RemoteRegistry, name string) (RemoteRegistryItem, error) {
	removedRemoteRegistry := RemoteRegistryItem{}

	index, err := getRegistryItemIndexByName(remoteRegistry, name)
	if err != nil {
		return removedRemoteRegistry, err
	}

	removedRemoteRegistry = remoteRegistry[index]

	remoteRegistry = slices.Delete(remoteRegistry, index, 1)
	err = writeConfig(remoteRegistry)
	if err != nil {
		return removedRemoteRegistry, err
	}

	return removedRemoteRegistry, nil

}

func ListRemoteRegistryIndex(remoteRegistry RemoteRegistry, remoteName string) (index.RegistryIndex, error) {
	registryIndex := index.RegistryIndex{}

	remoteRegistryItem, err := getRegistryItemByName(remoteRegistry, remoteName)
	if err != nil {
		return registryIndex, err
	}

	err = fetchRegistryIndex(remoteRegistryItem.Url, &registryIndex)
	if err != nil {
		return registryIndex, err
	}

	return registryIndex, nil
}

func PullFromRemoteRegistry(remoteRegistry RemoteRegistry, remoteName string, registryName string, outDir string, force bool) error {
	remoteRegistryItem, err := getRegistryItemByName(remoteRegistry, remoteName)
	if err != nil {
		return err
	}

	registryIndex := index.RegistryIndex{}
	err = fetchRegistryIndex(remoteRegistryItem.Url, &registryIndex)
	if err != nil {
		return err
	}

	registryIndexItem, err := index.FindRegistryItemByName(registryIndex, registryName)
	if err != nil {
		return err
	}

	files := []index.RegistryFile{}
	err = fetchRegistryItem(remoteRegistryItem.Url, registryIndexItem.Filename, &files)
	if err != nil {
		return err
	}

	outDirAbsPath, err := filepath.Abs(outDir)
	if err != nil {
		return err
	}

	// TODO each files rewrite file
	err = writeRegistryFiles(files, outDirAbsPath, force)
	if err != nil {
		return err
	}

	return nil
}
