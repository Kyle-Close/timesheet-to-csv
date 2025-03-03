package main

import (
	"log"
	"os"
)

func OpenFile(fileName string) []byte {
	// Get the filename that was passed in as arg
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
