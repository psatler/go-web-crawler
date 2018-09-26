package main

import (
	// import standard libraries
	"fmt"
	"log"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
)

//##### global stuff ####
var baseURL = "https://www.fundamentus.com.br/"

// var urlLinkSlices [887]string //mudar isso dps
// var urlSlices = make([]string, 887)

var allUrls []string

type PapersInfo struct {
	paperName, companyName string
	marketValue, dailyRate float64
}

var papersSlice = make([]PapersInfo, 887)

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

func NumberOfElementChild(s *goquery.Selection) int {
	return s.Children().Length()
	//return s.Children().Size()
}

func getInfoFromURL() {
	var paperName string = ""
	var companyName string = ""
	var marketValue string = ""
	var dailyRate string = ""
	// count := 0

	for i := 0; i < len(allUrls); i++ {
		doc, err := goquery.NewDocument(baseURL + allUrls[i])
		fmt.Println("\n" + baseURL + allUrls[i])
		if err != nil {
			log.Fatal(err)
		}

		// println(NumberOfElementChild(doc.Find("table.w728 tbody td.data")))

		//:nth-child(n)	p:nth-child(2)	Selects every <p> element that is the second child of its parent
		doc.Find("td.data.w35").Each(func(index int, item *goquery.Selection) { //company's name
			paperName = item.Find("span").Text()
			fmt.Printf("paper Name #%d: %s\n", index, paperName)
		})
		doc.Find("tr:nth-child(3) td:nth-child(2)").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			companyName = item.Find("span").Text()
			fmt.Printf("company Name #%d: %s\n", index, companyName)
			if index == 0 {
				return false
			}
			return true
		})
		doc.Find("td:nth-child(2) ").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			if index == 5 {
				marketValue = item.Find("span").Text()
				fmt.Printf("market Value #%d: %s\n", index, marketValue)
				return false
			}
			return true
		})

		doc.Find("span.oscil").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			if index == 0 {
				dailyRate = item.Text()
				fmt.Printf("market Value #%d: %s\n", index, dailyRate)
				return false
			}
			return true
		})
	}
}

func main() {

	getPaperLinks()
	fmt.Println(len(allUrls))
	getInfoFromURL()
}

// doc.Find("td.data.w35").Each(func(index int, item *goquery.Selection) { //company's name
// 	paperName := item.Find("span").Text()
// 	fmt.Printf("paper Name #%d: %s\n", index, paperName)
// })

// doc.Find("table.w728 tr td.data").Each(func(index int, item *goquery.Selection) { //company's name
// 	if index == 4 {
// 		companyName := item.Find("span").Text()
// 		fmt.Printf("company Name #%d: %s\n", index, companyName)
// 	}
// })

// doc.Find("tr td.data.w3").Each(func(index int, item *goquery.Selection) { //company's name
// 	if index == 1 {
// 		valorDeMercado := item.Find("span").Text()
// 		fmt.Printf("Valor de Mercado #%d: %s\n", index, valorDeMercado) //eh o segundo
// 	}
// })
