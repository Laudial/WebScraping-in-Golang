package main

import (
	"fmt"

	"WebScraping/controllers"
	"WebScraping/utils"
)

func main() {

	var (
		link string
	)

	// Nettoyer le terlminale
	utils.ClearScreen()

	fmt.Print("Site à scraper : ")
    fmt.Scan(&link)

	controllers.StartScraping(link)
}