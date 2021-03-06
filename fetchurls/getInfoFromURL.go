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
		paperInfo := globals.PapersInfo{} //declaring a struct of paper with its information

		response, err := http.Get(globals.BaseURL + globals.AllUrls[i])
		checkError(err, "at Get "+globals.BaseURL+globals.AllUrls[i])

		defer response.Body.Close()
		doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
		checkError(err, "at NewDocFromReader "+globals.BaseURL+globals.AllUrls[i])

		pNameSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(2) > tbody > tr:nth-child(1) > td.data.w35"
		// doc.Find("td.data.w35").Each(func(index int, item *goquery.Selection) { //company's name
		doc.Find(pNameSelector).Each(func(index int, item *goquery.Selection) { //company's name
			paperName = item.Find("span").Text()
			// fmt.Printf("paper Name #%d: %s\n", index, paperName)
			paperInfo.PaperName = paperName
		})

		cNameSelector := "body > div.center > div.conteudo.clearfix > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2)"
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

		//using a mutex to avoid losing data
		globals.AllPapersInfoStruct.Mu.Lock()
		globals.AllPapersInfoStruct.AllPapersInfo = append(globals.AllPapersInfoStruct.AllPapersInfo, paperInfo)
		globals.AllPapersInfoStruct.Mu.Unlock()

	}
}
