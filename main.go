package main

import (
	"fmt"
	"strings"

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
    _, err :=fmt.Scan(&link)

	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée :", err)
		return
	}

	if !strings.HasPrefix(link, "https://") {
		link = "https://" + link
	}

	controllers.StartScraping(link)
}