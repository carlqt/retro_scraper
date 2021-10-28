package main

import (
	"time"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0")
	time.Sleep(10 * time.Second)
	page.MustWaitLoad().MustScreenshot("a.png")
}
