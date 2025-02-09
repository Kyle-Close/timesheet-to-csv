package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gocarina/gocsv"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Too few arguments.")
	}

	file := CreateCSV()

	for i := 1; i < len(os.Args); i++ {
		data := OpenFile(os.Args[i])

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
