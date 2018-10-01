package utilfuncs

import (
	"sort"

	"github.com/psatler/go-web-crawler/globals"
)

// type stockInfo = globals.PapersInfo

func SortStockPapers() {
	sort.Slice(globals.AllPapersInfoStruct.AllPapersInfo, func(i, j int) bool {
		return globals.AllPapersInfoStruct.AllPapersInfo[i].MarketValue > globals.AllPapersInfoStruct.AllPapersInfo[j].MarketValue
	})
}

// func sortPapers2(info []stockInfo) {
// 	sort.Slice(info.AllPapersInfo, func(i, j int) bool {
// 		return info.AllPapersInfo[i].MarketValue > info.AllPapersInfo[j].MarketValue
// 	})
// }
