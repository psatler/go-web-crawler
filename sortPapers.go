package main

import "sort"

func sortPapers() {
	sort.Slice(allPapersInfo, func(i, j int) bool {
		return allPapersInfo[i].marketValue > allPapersInfo[j].marketValue
	})
}

