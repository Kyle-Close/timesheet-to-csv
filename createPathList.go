package main

import (
	"log"
	"os"
	"path"
	"time"
)

func createPathList(selectedWeek time.Time) []string {
	// selectedWeek will always be the start of the week (sunday)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	var paths []string
	directory := path.Join(homeDir, "Documents", "Mee", "work", "timesheets")

	for i := 0; i < 5; i++ {
		fileName := "Timesheet_" + selectedWeek.AddDate(0, 0, i+1).Format(time.DateOnly) + ".md"
		path := path.Join(directory, fileName)
		paths = append(paths, path)
	}

	return paths
}
