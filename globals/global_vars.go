package globals

import "sync"

//##### global stuff ####
var BaseURL = "https://www.fundamentus.com.br/"

var AllUrls []string

//it stores the information retrieved from the details page
type PapersInfo struct {
	PaperName, CompanyName, DailyRate string
	MarketValue                       float64
}

// var allPapersInfo []PapersInfo
var AllPapersInfoStruct struct {
	Mu            sync.Mutex
	AllPapersInfo []PapersInfo
}

var Wg sync.WaitGroup
