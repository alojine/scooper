package writer

import (
	"log"
	"os"
	"time"
)

const defaultFilePermissions = 0644

func WriteDataToFile(fileName string, data []byte) {
	if fileName == "" {
		log.Fatal("no file name was provided")
	}

	currentTime := time.Now().String()
	completeFileName := fileName + "_" + currentTime

	if data == nil {
		log.Fatal("no data was provided")
	}
	err := os.WriteFile(completeFileName, data, defaultFilePermissions)
	if err != nil {
		log.Fatal(err)
	}
}
