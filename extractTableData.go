package main

import (
	"log"
	"strconv"
	"strings"
)

type RowData struct {
	Name  string  `csv:"name"`
	Hours float32 `csv:"hours"`
	Date  string  `csv:"date"`
	Link  string  `csv:"link"`
	Note  string  `csv:"note"`
}

func ExtractTableData(tableSlice []string, date string) []RowData {
	rowsData := make([]RowData, 0, 50)

	for _, v := range tableSlice {
		parts := strings.Split(strings.TrimSpace(v), "|")
		if len(parts) < 5 {
			log.Fatal("Table has malformed data.")
		}

		// find index of start and end to grab number
		start := strings.IndexRune(parts[1], ':')
		end := strings.IndexRune(parts[1], ')')

		if start == -1 || end == -1 {
			continue
		}

		hoursStr := strings.TrimSpace(parts[1][start+3 : end])
		if hoursStr == "" {
			continue
		}

		hours, err := strconv.ParseFloat(hoursStr, 32)
		if err != nil {
			continue
		}
		name := strings.TrimSpace(parts[2])
		link := strings.TrimSpace(parts[3])
		note := strings.TrimSpace(parts[4])

		// Grab just the link not the full markdown [](link)
		start = strings.IndexRune(link, '(')
		end = strings.IndexRune(link, ')')

		if start != -1 && end != -1 {
			link = link[start+1 : end]
		}

		row := RowData{
			Hours: float32(hours),
			Name:  name,
			Date:  date,
			Link:  link,
			Note:  note,
		}

		rowsData = append(rowsData, row)
	}

	return rowsData
}
