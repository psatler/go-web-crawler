package main

import (
	"fmt"
	"strconv"
)

func printFirst10Papers() {
	for i := 0; i < 10; i++ {
		mValueInFloat := strconv.FormatFloat(allPapersInfoStruct.allPapersInfo[i].marketValue, 'f', 6, 64)
		fmt.Printf("\n#%d - \t Company: %s \n \t Market Value: %s \n", i, allPapersInfoStruct.allPapersInfo[i].companyName, mValueInFloat)

	}
}
