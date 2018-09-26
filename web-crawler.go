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
var urlLinkSlices [887]string //mudar isso dps
var urlSlices = make([]string, 887)

type PapersInfo struct {
	paperName, companyName string
	marketValue, dailyRate float64
}

var papersSlice = make([]PapersInfo, 887)

//find all the links in the main page
func paperScrape() {
	doc, err := goquery.NewDocument(baseURL + "detalhes.php")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tbody td a").Each(func(index int, item *goquery.Selection) { //using the HTML tag as selectors
		title := item.Text()
		link, _ := item.Attr("href") //get the link itself
		// linkTag := item.Find("a")
		// link, _ := linkTag.Attr("href")
		// fmt.Printf("Post #%d: %s - %s\n", index, title, link)
		urlLinkSlices[index] = link
		urlSlices[index] = link
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

	for i := 0; i < len(urlSlices); i++ {
		doc, err := goquery.NewDocument(baseURL + urlSlices[i])
		fmt.Println("\n" + baseURL + urlSlices[i])
		if err != nil {
			log.Fatal(err)
		}

		println(NumberOfElementChild(doc.Find("table.w728 tbody td.data")))

		//todos em um só
		doc.Find("table.w728 tr td.data").Each(func(index int, item *goquery.Selection) { //company's name
			if index == 0 {
				paperName = item.Find("span").Text()
				fmt.Printf("paper Name #%d: %s\n", index, paperName)
			}
			if index == 4 {
				companyName = item.Find("span").Text()
				fmt.Printf("company Name #%d: %s\n", index, companyName)
			}
			if index == 10 {
				marketValue = item.Find("span").Text()
				// marketV, _ := strconv.ParseFloat(marketValueString, 64)
				fmt.Printf("Valor de Mercado #%d: %s\n", index, marketValue) //eh o segundo
			}
			if index == 14 {
				dailyRate = item.Find("span").Text()
				// dailyR, _ := strconv.ParseFloat(dailyRateString, 64)
				fmt.Printf("Porcentagem Oscilacao Diaria #%d: %s\n", index, dailyRate) //eh o segundo
			}
		})

	}
}

func main() {
	paperScrape()
	fmt.Println(len(urlLinkSlices))
	fmt.Println(len(urlSlices))
	// for i := 0; i < len(urlSlices); i++ {
	// 	fmt.Println(urlSlices[i])
	// }
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
