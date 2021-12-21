package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/keenfury/api/config"
	ae "github.com/keenfury/api/internal/api_error"
)

func createFile(filePathWithName string) {
	if config.StorageFile {
		if _, err := os.Stat(filePathWithName); os.IsNotExist(err) {
			if err := os.WriteFile(filePathWithName, []byte("[]"), 0644); err != nil {
				fmt.Println("error creating storage folder:", err)
			}
		}
	}
}

func OpenFile(name string, obj interface{}) error {
	filePathWithName := fmt.Sprintf("%s/%s", config.StorageFilePath, name)
	createFile(filePathWithName)
	content, err := os.ReadFile(filePathWithName)
	if err != nil {
		return ae.GeneralError("unable to open file", err)
	}
	if err := json.Unmarshal(content, obj); err != nil {
		return ae.GeneralError("unable to decode file", err)
	}
	return nil
}

func SaveFile(name string, obj interface{}) error {
	filePathWithName := fmt.Sprintf("%s/%s", config.StorageFilePath, name)
	createFile(filePathWithName)
	content, err := json.Marshal(obj)
	if err != nil {
		return ae.GeneralError("unable to encode file", err)
	}
	if err := os.WriteFile(filePathWithName, content, 0644); err != nil {
		return ae.GeneralError("unable to save file", err)
	}
	return nil
}
