package index

import (
	"fmt"
	"glitchkids/registry_gk/helpers"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func withIsValidPath(output string, ignores ConfigInputFolderIgnore) func(path string, d fs.DirEntry) bool {
	return func(path string, d fs.DirEntry) bool {
		if d.IsDir() || strings.Contains(path, string(output)) {
			return false
		}

		name := d.Name()
		for _, v := range ignores.Files {
			if strings.Contains(v, "*") {
				if match, err := filepath.Match(v, name); match == true {
					return false
				} else if err != nil {
					log.Fatal(err)
				}
			}

			if path == v {
				return false
			}
		}

		for _, v := range ignores.Folders {
			if strings.Contains(path, fmt.Sprintf("/%v/", v)) {
				return false
			}
		}

		return true
	}
}

func walkRegisterFolder(configProcessPayload *RegistryFolder, isValidPath func(path string, d fs.DirEntry) bool, absPath string) func(path string, d fs.DirEntry, err error) error {
	return func(path string, d fs.DirEntry, err error) error {
		if d.Name() == CONFIG_FILENAME || isValidPath(path, d) == false {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, p, _ := strings.Cut(path, absPath)
		current := RegistryFile{Content: string(content), Path: fmt.Sprintf(".%v", p)}
		configProcessPayload.Files = append(configProcessPayload.Files, current)

		return nil
	}
}

func withWriteFile(outputPath string) func(name string, content []byte) error {
	return func(name string, content []byte) error {
		path := filepath.Join(outputPath, name)
		return helpers.WriteFile(path, content)
	}
}

func getFilename(c RegistryFolder) string {
	filename := strings.ReplaceAll(strings.ToLower(c.Name), "/", "_")
	return fmt.Sprintf("%v.json", filename)
}
