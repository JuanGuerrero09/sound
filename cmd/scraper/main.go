package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sound/pkg"

	"github.com/gocolly/colly"
)

func main() {
	pkg.Scraper()
	fmt.Println("Hola mundo")
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("could not create te file, err: %q", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("aliexpress.com"),
	)

}