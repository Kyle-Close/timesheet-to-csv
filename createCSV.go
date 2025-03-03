package main

import (
	"log"
	"os"
)

func CreateCSV() *os.File {
	file, err := os.Create("/home/kyle/Documents/timesheet-output.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	return file
}
