package writer

import (
	"fmt"
	"log"
	"os"
	"time"
)

const defaultFilePermissions = 0644
const defaultStoragePath = "./storage/"

func WriteDataToFile(fileName string, data []byte) {
	if fileName == "" {
		log.Fatal("no file name was provided")
	}

	if data == nil {
		log.Fatal("no data was provided")
	}

	if _, err := os.Stat(defaultStoragePath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(defaultStoragePath, defaultFilePermissions)
		}
	}

	storageFilePath := getStorageFilePath(fileName)
	err := os.WriteFile(storageFilePath, data, defaultFilePermissions)
	if err != nil {
		log.Fatal(err)
	}
}

func getStorageFilePath(fileName string) string {
	t := time.Now()
	formattedTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	completeFilePath := defaultStoragePath + fileName + "_" + formattedTime
	return completeFilePath
}
