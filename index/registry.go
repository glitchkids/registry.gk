package index

import (
	"fmt"
	"glitchkids/registry_gk/helpers"
	"os"
	"path/filepath"
	"slices"
)

type Registry struct {
	Index   RegistryIndex
	Folders []RegistryFolder
	Files   []RegistryFile
}

type RegistryIndex []RegistryIndexItem

func (r RegistryIndex) String() string {
	var s string
	for i, v := range r {
		s = fmt.Sprintf("%v%v. %v", s, i, v)
	}
	return s
}

type RegistryIndexItem struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

func (r RegistryIndexItem) String() string {
	return fmt.Sprintf("%v\n", r.Name)
}

type RegistryFolder struct {
	Name  string
	Files []RegistryFile
}

type RegistryFile struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func Save(c Config, registry Registry) error {
	_, err := os.Stat(c.Output)
	if err == nil {
		os.RemoveAll(c.Output)
	}

	err = helpers.Mkdir(c.Output)
	if err != nil {
		return err
	}

	writeFile := withWriteFile(c.Output)

	for _, f := range registry.Folders {
		content, err := helpers.ToJSON(f)
		if err != nil {
			return err
		}

		err = writeFile(getFilename(f), []byte(content))
		if err != nil {
			return err
		}
	}

	indexContent, err := helpers.ToJSON(registry.Index)
	if err != nil {
		return err
	}

	err = writeFile("index.json", []byte(indexContent))
	if err != nil {
		return err
	}

	return nil
}

func CreateRegistry(c Config) (Registry, error) {
	registry := Registry{}

	for _, f := range c.Input.Folders {
		isValidPath := withIsValidPath(c.Output, f.Ignores)

		if filepath.IsLocal(f.Path) == false {
			return registry, fmt.Errorf("Local path is required : %v", f.Name)
		}

		absPath, err := filepath.Abs(f.Path)
		if err != nil {
			return registry, err
		}

		if _, err = os.Stat(absPath); err != nil {
			return registry, err
		}

		registryFolder := RegistryFolder{Name: f.Name}
		err = filepath.WalkDir(absPath, walkRegisterFolder(&registryFolder, isValidPath, absPath))
		if err != nil {
			return registry, err
		}

		registry.Folders = append(registry.Folders, registryFolder)
	}

	index := []RegistryIndexItem{}
	for _, f := range registry.Folders {
		index = append(index, RegistryIndexItem{Name: f.Name, Filename: getFilename(f)})
	}
	registry.Index = index

	return registry, nil
}

func GetIndexItemByName(r RegistryIndex, name string) (RegistryIndexItem, error) {
	index := slices.IndexFunc(r, func(e RegistryIndexItem) bool {
		return e.Name == name
	})

	if index == -1 {
		return RegistryIndexItem{}, fmt.Errorf("Registry not found : %v", name)
	}

	return r[index], nil
}
