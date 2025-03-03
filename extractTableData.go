package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

type RowData struct {
	Hours        float32 `csv:"hours"`
	Name         string  `csv:"name"`
	TaskCategory string  `csv:"taskCategory"`
	Date         string  `csv:"date"`
	Link         string  `csv:"link"`
	Note         string  `csv:"note"`
}

func ExtractTableData(tableSlice []string, date string) []RowData {
	rowsData := make([]RowData, 0, 50)

	for _, v := range tableSlice {
		columns := strings.Split(strings.TrimSpace(v), "|")

		if len(columns) < getStructFieldCount[RowData]() {
			log.Fatal("Table has malformed data.")
		}

		// find index of start and end to grab number
		start := strings.IndexRune(columns[1], ':')
		end := strings.IndexRune(columns[1], ')')

		if start == -1 || end == -1 {
			continue
		}

		hoursStr := strings.TrimSpace(columns[1][start+3 : end])
		if hoursStr == "" {
			continue
		}

		hours, err := strconv.ParseFloat(hoursStr, 32)
		if err != nil {
			continue
		}

		name := strings.TrimSpace(columns[2])
		taskCategory := strings.TrimSpace(columns[3])
		link := strings.TrimSpace(columns[4])
		note := strings.TrimSpace(columns[5])

		// Grab just the link not the full markdown [](link)
		start = strings.IndexRune(link, '(')
		end = strings.IndexRune(link, ')')

		if start != -1 && end != -1 {
			link = link[start+1 : end]
		}

		row := RowData{
			Hours:        float32(hours),
			Name:         name,
			TaskCategory: taskCategory,
			Date:         date,
			Link:         link,
			Note:         note,
		}

		rowsData = append(rowsData, row)
	}

	return rowsData
}

func getStructFieldCount[T any]() int {
	return reflect.TypeOf((*T)(nil)).Elem().NumField()
}
