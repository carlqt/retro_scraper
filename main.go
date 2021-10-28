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
	page.MustWaitElementsMoreThan(".column-header", 1)

	headers := page.MustWaitLoad().MustElements(".column-header")

	for _, header := range headers {
		value := header.MustElement("h2.no-edit").MustText()

		fmt.Println(value)
	}
}
