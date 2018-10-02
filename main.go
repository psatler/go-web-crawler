package main

import (
	//custom packages
	"github.com/psatler/go-web-crawler/databaseutils"
	"github.com/psatler/go-web-crawler/fetchurls"
	"github.com/psatler/go-web-crawler/globals"
	"github.com/psatler/go-web-crawler/utilfuncs"

	// import standard libraries
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())          //4
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //4

	fetchurls.GetPaperLinks()
	fmt.Println(len(globals.AllUrls))
	size := len(globals.AllUrls)
	// divisor := 6 //for some reason it's been optimal for me in my computer
	// divisor := 6 //for some reason it's been optimal for me in my computer and it takes 3m52.811s
	// divisor := 30 //54s
	// divisor := 50 //3m30.269s
	divisor := 80 //2m30.975s
	// divisor := 85 //3m40.942s
	// divisor := 88 //3m11.998s

	globals.Wg.Add(divisor)
	for i := 0; i < divisor; i++ {
		if (divisor - i) == 1 { //if is the last iteration, take care of summing the remainder
			go fetchurls.GetInfoFromURL((size/divisor)*i, (size/divisor)*(i+1)+(size%divisor))
		} else {
			go fetchurls.GetInfoFromURL((size/divisor)*i, (size/divisor)*(i+1))
		}
	}
	globals.Wg.Wait()

	fmt.Printf("allPapersInfoStruct: %d ", len(globals.AllPapersInfoStruct.AllPapersInfo))
	utilfuncs.FirstTen()
	utilfuncs.SortStockPapers()
	fmt.Println("\n - In Order ")
	utilfuncs.FirstTen()

	fmt.Println("\n WRITING TO DB -")
	databaseutils.WriteToDb()
	fmt.Println("\n READING FROM DB -")
	result := databaseutils.ReadFromDb()
	fmt.Println("\n - PRINTING FROM DATABASE ")
	for i := 0; i < len(result); i++ {
		fmt.Printf("#%d - \t Company: %s \n \t Market Value: %.2f \n", i,
			result[i].CompanyName, result[i].MarketValue)
	}

}
