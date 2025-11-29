package writer

import (
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

	currentTime := time.Now().String()
	completeFilePath := defaultStoragePath + fileName + "_" + currentTime

	if data == nil {
		log.Fatal("no data was provided")
	}

	if _, err := os.Stat(defaultStoragePath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(defaultStoragePath, defaultFilePermissions)
		}
	}

	err := os.WriteFile(completeFilePath, data, defaultFilePermissions)
	if err != nil {
		log.Fatal(err)
	}
}
