package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Too few arguments.")
	}

	file := CreateCSV()

	for i := 1; i < len(os.Args); i++ {
		data := OpenFile(os.Args[i])
		tableSlice := ExtractTable(data)
		tableData := ExtractTableData(tableSlice)
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
