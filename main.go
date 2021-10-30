package main

import (
	"encoding/csv"
	"os"

	"github.com/go-rod/rod"
)

type Column struct {
	Header   string
	Contents []string
}

func NewColumns(driver *rod.Page) []Column {
	var columns []Column
	columnsElement := driver.MustElements(".easy-card-list")

	for _, element := range columnsElement {
		contents := make([]string, 1)
		header := element.MustElement("h2.no-edit").MustText()

		columnContents := element.MustElements("ul.column > li")
		for _, contentElement := range columnContents {
			text := contentElement.MustElement(".easy-board-front .easy-card-main-content .text").MustText()
			contents = append(contents, text)
		}

		col := Column{
			Header:   header,
			Contents: contents,
		}
		columns = append(columns, col)
	}

	return columns
}

func exportToCSV(columns []Column) {
	// convert columns into something
	inputs := convertToCsvInput(columns)

	// maybe add timestamp
	os.MkdirAll("output", os.ModePerm)
	file, err := os.Create("output/easyretro.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, input := range inputs {
		_ = writer.Write(input)
	}
}

func convertToCsvInput(columns []Column) [][]string {
	// detect dimensions
	// detect number of max rows
	var rowsCount int
	for _, column := range columns {
		if rowsCount < len(column.Contents) {
			rowsCount = len(column.Contents)
		}
	}

	// detect number of columns
	columnsCount := len(columns)

	// initialize the board
	cells := make([][]string, rowsCount+1)
	for i := range cells {
		cells[i] = make([]string, columnsCount)
	}

	for y, column := range columns {
		cells[0][y] = column.Header

		for x, content := range column.Contents {
			cells[x+1][y] = content
		}
	}

	return cells
}

func main() {
	// Test URL - https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0"
	argsURL := os.Args[1]

	page := rod.New().MustConnect().MustPage(argsURL)
	page.MustWaitElementsMoreThan(".easy-card-list", 0)

	columns := NewColumns(page)

	exportToCSV(columns)
}
