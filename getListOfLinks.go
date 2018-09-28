package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

//find all the links in the main page
func getPaperLinks() {
	doc, err := goquery.NewDocument(baseURL + "detalhes.php")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tbody td a").Each(func(index int, item *goquery.Selection) { //using the HTML tag as selectors
		title := item.Text()
		link, _ := item.Attr("href") //get the link itself
		allUrls = append(allUrls, link)
		fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
}
