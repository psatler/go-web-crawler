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
	paperName, companyName, dailyRate string
	marketValue                       float64
}

var allPapersInfo []PapersInfo
var wg sync.WaitGroup

func main() {
	// var wg sync.WaitGroup
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())          //4
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //4

	// runtime.GOMAXPROCS(runtime.NumCPU())

	getPaperLinks()
	fmt.Println(len(allUrls))
	size := len(allUrls)
	// divisor := 6 //for some reason it's been optimal for me in my computer
	// divisor := 6 //for some reason it's been optimal for me in my computer
	// divisor := 30 //54s
	// divisor := 50 //0m45.156s
	divisor := 80 //0m26.239s
	// divisor := 200 //0m26.239s

	// wg.Add(divisor)
	// go getInfoFromURL((size / divisor * 0), size/divisor*1)
	// go getInfoFromURL((size / divisor * 1), size/divisor*2)
	// go getInfoFromURL((size / divisor * 2), size/divisor*3+(size%divisor))
	// wg.Wait()

	wg.Add(divisor)
	for i := 0; i < divisor; i++ {
		if (divisor - i) == 1 { //if is the last iteration, take care of summing the remainder
			go GetInfoFromURL((size/divisor)*i, (size/divisor)*(i+1)+(size%divisor))
		} else {
			go GetInfoFromURL((size/divisor)*i, (size/divisor)*(i+1))
		}
	}
	wg.Wait()

	fmt.Printf("allPapersInfo: %d ", len(allPapersInfo))
	printFirst10Papers()
	sortPapers()
	fmt.Println("\n - In Order - \n")
	printFirst10Papers()

	fmt.Println("\n teste 1 -")
	WriteToDb()
	fmt.Println("\n teste 2 -")
	result := ReadFromDb()
	for i := 0; i < len(result); i++ {
		fmt.Println("\n - FROM DATABASE - \n")
		fmt.Printf("\n#%d - \t Company: %s \n \t Market Value: %f \n", i, allPapersInfo[i].companyName, allPapersInfo[i].marketValue)
	}
}
