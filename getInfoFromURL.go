package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetInfoFromURL(init int, end int) {
	var paperName string
	var companyName string
	var marketValue float64
	var dailyRate string

	// size := len(allUrls)
	defer wg.Done()

	for i := init; i < end; i++ {
		// for i := 0; i < len(allUrls); i++ {
		paperInfo := PapersInfo{} //a struct of paper with its information
		doc, err := goquery.NewDocument(baseURL + allUrls[i])
		// fmt.Println(i, " - ", baseURL+allUrls[i])
		if err != nil {
			fmt.Println(i, " - ", baseURL+allUrls[i])
			log.Fatal("GetInfoFromURL.go file: ", err)
		}

		// println(NumberOfElementChild(doc.Find("table.w728 tbody td.data")))

		//https://www.w3schools.com/cssref/css_selectors.asp
		pNameSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(2) > tbody > tr:nth-child(1) > td.data.w35"
		// doc.Find("td.data.w35").Each(func(index int, item *goquery.Selection) { //company's name
		doc.Find(pNameSelector).Each(func(index int, item *goquery.Selection) { //company's name
			paperName = item.Find("span").Text()
			// fmt.Printf("paper Name #%d: %s\n", index, paperName)
			paperInfo.paperName = paperName
		})
		cNameSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2)"
		// doc.Find("tr:nth-child(3) td:nth-child(2)").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
		doc.Find(cNameSelector).EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			companyName = item.Find("span").Text()
			paperInfo.companyName = companyName
			// fmt.Printf("company Name #%d: %s\n", index, companyName)
			if index == 0 {
				return false
			}
			return true
		})
		mvalueSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(3) > tbody > tr:nth-child(1) > td.data.w3"
		doc.Find(mvalueSelector).EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			if index == 0 {
				marketV := item.Find("span").Text()             //text as string
				noDots := strings.Replace(marketV, ".", "", -1) //-1 means all occurrencies (taking out the dots in the string to convert it to float later)
				marketValue, _ = strconv.ParseFloat(noDots, 64) //converting to float
				paperInfo.marketValue = marketValue
				// fmt.Println(marketValue)
				// fmt.Println(strconv.FormatFloat(marketValue, 'f', 6, 64))
				// fmt.Printf("market Value #%d: %f\n", index, marketValue)

				return false
			}
			return true
		})

		// dRateSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(4) > tbody > tr:nth-child(2) > td.data.w1 > span > font"
		doc.Find("span.oscil").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			if index == 0 {
				dailyRate = item.Text()
				// fmt.Printf("daily Rate #%d: %s\n", index, dailyRate)
				paperInfo.dailyRate = dailyRate
				return false
			}
			return true
		})

		//append the newly created struct to the slice of all papers
		allPapersInfo = append(allPapersInfo, paperInfo)

		// fmt.Println(len(allPapersInfo))
	}
}
