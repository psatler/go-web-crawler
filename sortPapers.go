package main

import "sort"

func sortPapers() {
	sort.Slice(allPapersInfoStruct.allPapersInfo, func(i, j int) bool {
		return allPapersInfoStruct.allPapersInfo[i].marketValue > allPapersInfoStruct.allPapersInfo[j].marketValue
	})
}
