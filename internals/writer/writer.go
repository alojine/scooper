package writer

import (
	"fmt"
	"os"
	"time"
)

const defaultFilePermissions = 0644
const defaultStoragePath = "./storage/"

func WriteDataToFile(fileName string, data []byte) error {
	if fileName == "" {
		return fmt.Errorf("no file name was provided")
	}

	if data == nil {
		return fmt.Errorf("no data was provided")
	}

	if err := os.MkdirAll(defaultStoragePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory for file storing: %w", err)
	}

	storageFilePath := getStorageFilePath(fileName)
	if err := os.WriteFile(storageFilePath, data, defaultFilePermissions); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func getStorageFilePath(fileName string) string {
	t := time.Now()
	formattedTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	completeFilePath := defaultStoragePath + fileName + "_" + formattedTime
	return completeFilePath
}
