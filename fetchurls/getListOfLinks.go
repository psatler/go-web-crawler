package fetchurls

import (
	"log"

	"github.com/PuerkitoBio/goquery"

	//custom package
	"github.com/psatler/go-web-crawler/globals"
)

var baseURL = globals.BaseURL

//find all the links in the main page
func GetPaperLinks() {
	doc, err := goquery.NewDocument(baseURL + "detalhes.php")
	if err != nil {
		log.Fatal("getPaperLinks", err)
	}

	doc.Find("tbody td a").Each(func(index int, item *goquery.Selection) { //using the HTML tag as selectors
		// title := item.Text()
		link, _ := item.Attr("href") //get the link itself
		globals.AllUrls = append(globals.AllUrls, link)
		// allUrls = append(allUrls, link)
		// fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
}
