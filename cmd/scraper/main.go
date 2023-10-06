package main

import (
	"fmt"
	// "encoding/csv"
	// "log"
	// "os"
	// "sound/pkg"

	"github.com/gocolly/colly"
)

func main() {
	// pkg.Scraper()
	// fmt.Println("Hola mundo")
	// fName := "data.csv"
	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatalf("could not create te file, err: %q", err)
	// }
	// defer file.Close()

	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// c := colly.NewCollector(
	// 	colly.AllowedDomains("http://quotes.toscrape.com"),
	// )

	// c.OnRequest(func(r *colly.Request) {
	// 	 r.Headers.Set("User-Agent")
	// 	fmt.Println("Visiting: ", r.URL)
	// })

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Response code: ", r.StatusCode)
	// })

	// c.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("error", err.Error())
	// })

	
	c := colly.NewCollector()
	
	// Find and visit all links
	c.OnHTML(".text", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	
	c.OnHTML(".author", func(h *colly.HTMLElement) {
		fmt.Println("Author: ", h.Text)
	})
	c.Visit("https://es.aliexpress.com/item/1005005816349636.html")



}