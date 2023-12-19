package controllers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func StartScraping(link string) {
	// Créez un nouveau collecteur
	c := colly.NewCollector()

	// Définissez les sélecteurs pour extraire les données
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		fmt.Println("Title:", title)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link:", link)
	})

	// Lancez le collecteur
	c.Visit(link)
}