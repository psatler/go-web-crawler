package main

import (
	"sync"

	// import standard libraries
	"fmt"
	"runtime"
	// import third party libraries
)

//##### global stuff ####
var baseURL = "https://www.fundamentus.com.br/"

var allUrls []string

type PapersInfo struct {
	paperName, companyName string
	marketValue, dailyRate float64
}

var allPapersInfo []PapersInfo
var wg sync.WaitGroup

// func NumberOfElementChild(s *goquery.Selection) int {
// 	return s.Children().Length()
// 	//return s.Children().Size()
// }

func main() {
	// var wg sync.WaitGroup
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())          //4
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //4

	getPaperLinks()
	fmt.Println(len(allUrls))
	size := len(allUrls)
	divisor := 6

	// wg.Add(divisor)
	// go getInfoFromURL((size / divisor * 0), size/divisor*1)
	// go getInfoFromURL((size / divisor * 1), size/divisor*2)
	// go getInfoFromURL((size / divisor * 2), size/divisor*3+(size%divisor))
	// wg.Wait()

	wg.Add(divisor)
	for i := 0; i < divisor; i++ {
		if (divisor - i) == 1 { //if is the last iteration, take care of summing the remainder
			go GetInfoFromURL((size / divisor * i), size/divisor*(i+1)+(size%divisor))
		} else {
			go GetInfoFromURL((size / divisor * i), size/divisor*(i+1))
		}
	}
	wg.Wait()

	fmt.Printf("allPapersInfo: %d ", len(allPapersInfo))
	printFirst10Papers()
	sortPapers()
	printFirst10Papers()
}
