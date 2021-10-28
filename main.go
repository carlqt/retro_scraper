package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

type Column struct {
	Header   string
	Contents []string
}

func NewColumns(driver *rod.Page) []Column {
	var columns []Column
	var contents []string
	columnsElement := driver.MustElements(".easy-card-list")

	for _, element := range columnsElement {
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

func main() {
	page := rod.New().MustConnect().MustPage("https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0")
	page.MustWaitElementsMoreThan(".easy-card-list", 0)

	columns := NewColumns(page)

	fmt.Println(columns)
}
