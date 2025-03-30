package main

import (
	"log"
	"os"
	"path/filepath"
)

func CreateCSV() *os.File {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	filePath := filepath.Join(homeDir, "Documents", "timesheet-output.csv")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	return file
}
