package main

import (
	"log"
	"os"
)

func main() {
	data := OpenFile()
	tableSlice := ExtractTable(data)
	tableData := ExtractTableData(tableSlice)

	for _, v := range tableData {
		log.Println(v)
	}

	os.Exit(0)
}
