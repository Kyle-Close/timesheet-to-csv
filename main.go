package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gocarina/gocsv"
)

func main() {
	// Get the time (starting sunday) that we want to create the CSV for
	selectedWeek := getWeekSelection()

	// Get the csv files for that week
	paths := createPathList(selectedWeek)
	file := CreateCSV()

	for i := 0; i < len(paths); i++ {
		data := OpenFile(paths[i])

		// Grab the date from the file
		re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
		date := strings.Split(re.FindString(string(data)), " ")[0]

		tableSlice := ExtractTable(data)
		tableData := ExtractTableData(tableSlice, date)
		err := gocsv.MarshalFile(&tableData, file)
		if err != nil {
			log.Fatal(err.Error())
		}
		if _, err := file.WriteString("\n\n"); err != nil { // Two empty rows
			log.Fatal(err)
		}
	}

	os.Exit(0)
}
