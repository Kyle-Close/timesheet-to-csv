package main

import (
	"log"
	"os"
)

func OpenFile() []byte {
	// Get the filename that was passed in as arg
	if len(os.Args) <= 1 {
		log.Fatal("Invalid amount of args. Expected filename of file.")
	}

	fileName := os.Args[1]
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
