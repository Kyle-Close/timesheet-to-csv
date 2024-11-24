package main

import (
	"log"
	"os"
)

func OpenFile(fileName string) []byte {
	// Get the filename that was passed in as arg
	if len(os.Args) <= 1 {
		log.Fatal("Invalid amount of args. Expected filename of file.")
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
