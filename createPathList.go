package main

import (
	"path"
	"time"
)

func createPathList(selectedWeek time.Time) []string {
	// selectedWeek will always be the start of the week (sunday)

	var paths []string
	directory := "/home/kyle/Documents/Mee/work/timesheets"

	for i := 0; i < 5; i++ {
		fileName := "Timesheet_" + selectedWeek.AddDate(0, 0, i+1).Format(time.DateOnly) + ".md"
		path := path.Join(directory, fileName)
		paths = append(paths, path)
	}

	return paths
}
