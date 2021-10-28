package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	// u := launcher.New().
	// 	Delete("--headless").
	// 	MustLaunch()

	// page := rod.New().ControlURL(u).MustConnect().MustPage("https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0")
	// page.MustWaitLoad().MustScreenshot("a.png")

	page := rod.New().MustConnect().MustPage("https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0")
	page.MustWaitElementsMoreThan(".easy-card-list", 0)

	columns := page.MustWaitLoad().MustElements(".easy-card-list")

	for _, column := range columns {
		header := column.MustElement("h2.no-edit").MustText()

		fmt.Println(header)

		contents := column.MustElements("ul.column > li")

		// Listing out the column's contents
		for _, content := range contents {
			text := content.MustElement(".easy-board-front .easy-card-main-content .text").MustText()
			fmt.Println(text)
		}
	}
}
