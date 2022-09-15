package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {

	fName := "data_xs.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("File not created, err : %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org", "xoso.com.vn", "www.w3schools.com"),
	)
	//	xsmmb
	c.OnHTML("div[class=section-content]", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.ChildText("td:nth-child(1)"))
		})
	})

	c.Visit("https:///xo-so-mien-bac/xsmb-p1.html")

	// w3school
	//c.OnHTML(".ws-table-all", func(e *colly.HTMLElement) {
	//	e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
	//		writer.Write([]string{
	//			el.ChildText("td:nth-child(1)"),
	//			el.ChildText("td:nth-child(2)"),
	//			el.ChildText("td:nth-child(3)"),
	//		})
	//	})
	//})
	//c.Visit("https://www.w3schools.com/html/html_tables.asp")

	// display crawl link
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	c.Visit(e.Request.AbsoluteURL(link))
	//})
	//
	//c.Limit(&colly.LimitRule{
	//	DomainGlob:  "*",
	//	RandomDelay: 1 * time.Second,
	//})
	//
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("crawling", r.URL.String())
	//})
}
