package helpers

import (
	"bytes"
	"encoding/json"
	"os"
)

func Access(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0666)
}

func Mkdir(path string) error {
	return os.MkdirAll(path, 0777)
}

func ToJSON(v any) (string, error) {
	var out bytes.Buffer

	jsonPayload, err := json.Marshal(v)
	if err != nil {
		return out.String(), err
	}

	err = json.Indent(&out, []byte(jsonPayload), "", "\t")
	if err != nil {
		return out.String(), err
	}

	return out.String(), nil
}
