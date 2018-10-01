package utilfuncs

import (
	"fmt"
	"strconv"

	//custom package
	"github.com/psatler/go-web-crawler/globals"
)

func FirstTen() {
	for i := 0; i < 10; i++ {
		mValueInFloat := strconv.FormatFloat(globals.AllPapersInfoStruct.AllPapersInfo[i].MarketValue, 'f', 6, 64)
		fmt.Printf("\n#%d - \t Company: %s \n \t Market Value: %s \n", i, globals.AllPapersInfoStruct.AllPapersInfo[i].CompanyName, mValueInFloat)

	}
}
