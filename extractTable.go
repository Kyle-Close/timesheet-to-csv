package main

import (
	"log"
	"math"
	"strings"
)

func ExtractTable(data []byte) []string {
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		log.Fatal("Error splitting file. No lines extracted.")
	}

	tableSlice := make([]string, 0, 50)
	startIndex := math.MaxInt

	for index, value := range lines {
		if strings.Contains(strings.ToUpper(value), "HOURS") && strings.Contains(strings.ToUpper(value), "NAME") {
			startIndex = index + 2
		}

		if index >= startIndex {
			if strings.Contains(value, "|") {
				tableSlice = append(tableSlice, value)
			} else {
				break
			}
		}
	}

	return tableSlice
}
