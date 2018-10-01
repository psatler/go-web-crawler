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

// var allPapersInfo []PapersInfo
var allPapersInfoStruct struct {
	mu            sync.Mutex
	allPapersInfo []PapersInfo
}

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
	divisor := 6 //for some reason it's been optimal for me in my computer and it takes 3m52.811s
	// divisor := 30 //54s
	// divisor := 50 //3m30.269s
	// divisor := 80 //2m30.975s
	// divisor := 85 //3m40.942s
	// divisor := 88 //3m11.998s

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

	// fmt.Printf("allPapersInfo: %d \n", len(allPapersInfo))
	fmt.Printf("allPapersInfoStruct: %d ", len(allPapersInfoStruct.allPapersInfo))
	printFirst10Papers()
	sortPapers()
	fmt.Println("\n - In Order ")
	printFirst10Papers()

	fmt.Println("\n WRITING TO DB -")
	WriteToDb()
	fmt.Println("\n READING FROM DB -")
	result := ReadFromDb()
	fmt.Println("\n - PRINTING FROM DATABASE ")
	for i := 0; i < len(result); i++ {
		fmt.Printf("\n#%d - \t Company: %s \n \t Market Value: %f \n", i, allPapersInfoStruct.allPapersInfo[i].companyName, allPapersInfoStruct.allPapersInfo[i].marketValue)
	}

}
