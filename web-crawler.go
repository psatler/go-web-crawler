package main

import (
	// import standard libraries
	"fmt"
	"log"
	"runtime"
	"sort"
	"strconv"
	"strings"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
)

//##### global stuff ####
var baseURL = "https://www.fundamentus.com.br/"

var allUrls []string

type PapersInfo struct {
	paperName, companyName string
	marketValue, dailyRate float64
}

var allPapersInfo []PapersInfo

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

func sortPapers() {
	sort.Slice(allPapersInfo, func(i, j int) bool { return allPapersInfo[i].marketValue > allPapersInfo[j].marketValue })
}

func printFirst10Papers() {
	for i := 0; i < 10; i++ {
		mValueInFloat := strconv.FormatFloat(allPapersInfo[i].marketValue, 'f', 6, 64)
		fmt.Printf("#%d - \t Company: %s \n \t Market Value: %s \n", i, allPapersInfo[i].companyName, mValueInFloat)
	}
}

func getInfoFromURL() {
	var paperName string = ""
	var companyName string = ""
	var marketValue float64
	var dailyRate string

	for i := 0; i < len(allUrls); i++ {
		// for i := 0; i < 30; i++ {
		paperInfo := PapersInfo{} //a struct of paper with its information
		doc, err := goquery.NewDocument(baseURL + allUrls[i])
		fmt.Println(i, " - ", baseURL+allUrls[i])
		if err != nil {
			log.Fatal(err)
		}

		// println(NumberOfElementChild(doc.Find("table.w728 tbody td.data")))

		//https://www.w3schools.com/cssref/css_selectors.asp
		doc.Find("td.data.w35").Each(func(index int, item *goquery.Selection) { //company's name
			paperName = item.Find("span").Text()
			fmt.Printf("paper Name #%d: %s\n", index, paperName)
			paperInfo.paperName = paperName
		})
		doc.Find("tr:nth-child(3) td:nth-child(2)").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			companyName = item.Find("span").Text()
			paperInfo.companyName = companyName
			fmt.Printf("company Name #%d: %s\n", index, companyName)
			if index == 0 {
				return false
			}
			return true
		})
		doc.Find("td:nth-child(2) ").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			if index == 5 {
				marketV := item.Find("span").Text()
				noDots := strings.Replace(marketV, ".", "", -1) //-1 means all occurrencies
				marketValue, _ = strconv.ParseFloat(noDots, 64)
				paperInfo.marketValue = marketValue
				// fmt.Println(marketValue)
				// fmt.Println(strconv.FormatFloat(marketValue, 'f', 6, 64))
				fmt.Printf("market Value #%d: %f\n", index, marketValue)

				return false
			}
			return true
		})

		doc.Find("span.oscil").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			if index == 0 {
				dailyRate = item.Text()
				fmt.Printf("daily Rate #%d: %s\n", index, dailyRate)
				return false
			}
			return true
		})

		//append the newly created struct to the slice of all papers
		allPapersInfo = append(allPapersInfo, paperInfo)

		// fmt.Println(len(allPapersInfo))
	}
}

func main() {
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())          //4
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //4

	getPaperLinks()
	fmt.Println(len(allUrls))
	getInfoFromURL()
	fmt.Println(len(allPapersInfo))
	printFirst10Papers()
	sortPapers()
	printFirst10Papers()
}
