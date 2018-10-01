package fetchurls

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	//custom package
	"github.com/psatler/go-web-crawler/globals"
)

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		// debug.PrintStack() //so we don't lose the stack trace
	}
}

func checkError(err error, url string) {
	if err != nil {
		fmt.Println("URL: " + url)
		panic(err)
		// os.Exit(1)
	}
}

func GetInfoFromURL(init int, end int) {
	var paperName string
	var companyName string
	var marketValue float64
	var dailyRate string

	defer recoverFunc()
	defer globals.Wg.Done()

	for i := init; i < end; i++ {
		// for i := 0; i < len(allUrls); i++ {
		paperInfo := globals.PapersInfo{} //declaring a struct of paper with its information

		response, err := http.Get(globals.BaseURL + globals.AllUrls[i])
		checkError(err, globals.BaseURL+globals.AllUrls[i])

		defer response.Body.Close()
		// println(response.Body)
		doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
		checkError(err, globals.BaseURL+globals.AllUrls[i])

		// doc, err := goquery.NewDocument(baseURL + allUrls[i])
		// // fmt.Println(i, " - ", baseURL+allUrls[i])
		// if err != nil {
		// 	fmt.Println(i, " - ", baseURL+allUrls[i])
		// 	log.Fatal("GetInfoFromURL.go file: ", err)
		// }

		// println(NumberOfElementChild(doc.Find("table.w728 tbody td.data")))

		//https://www.w3schools.com/cssref/css_selectors.asp
		pNameSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(2) > tbody > tr:nth-child(1) > td.data.w35"
		// doc.Find("td.data.w35").Each(func(index int, item *goquery.Selection) { //company's name
		doc.Find(pNameSelector).Each(func(index int, item *goquery.Selection) { //company's name
			paperName = item.Find("span").Text()
			// fmt.Printf("paper Name #%d: %s\n", index, paperName)
			paperInfo.PaperName = paperName
		})
		cNameSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2)"
		// doc.Find("tr:nth-child(3) td:nth-child(2)").EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
		doc.Find(cNameSelector).EachWithBreak(func(index int, item *goquery.Selection) bool { //company's name
			companyName = item.Find("span").Text()
			paperInfo.CompanyName = companyName
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
				paperInfo.MarketValue = marketValue
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
				paperInfo.DailyRate = dailyRate
				return false
			}
			return true
		})

		//append the newly created struct to the slice of all papers
		// allPapersInfo = append(allPapersInfo, paperInfo)

		//using a mutex to avoid losing data
		globals.AllPapersInfoStruct.Mu.Lock()
		globals.AllPapersInfoStruct.AllPapersInfo = append(globals.AllPapersInfoStruct.AllPapersInfo, paperInfo)
		globals.AllPapersInfoStruct.Mu.Unlock()

	}
}
