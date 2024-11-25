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
		if isHeaderLine(value) {
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

func isHeaderLine(line string) bool {
	headerVals := [4]string{"HOURS", "NAME", "LINK", "NOTES"}

	for _, v := range headerVals {
		if !strings.Contains(strings.ToUpper(line), v) {
			return false
		}
	}

	return true
}
