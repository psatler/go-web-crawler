package databaseutils

import (
	"database/sql"
	"fmt"
	"strconv"

	//custom package
	"github.com/psatler/go-web-crawler/globals"
)

//writes the result to a mysql database
func WriteToDb() {
	//db requests
	dataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dataSourceName) //open a connection
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	for i := 0; i < 10; i++ {
		mValueInFloat := strconv.FormatFloat(globals.AllPapersInfoStruct.AllPapersInfo[i].MarketValue, 'f', 0, 64)

		stmt, err := db.Prepare("insert into ibovespa (paperName, companyName, dailyRate, marketValue) values(?,?,?,?);")
		if err != nil {
			fmt.Print("At writing to db (Prepare): ", err.Error())
		}
		_, err = stmt.Exec(globals.AllPapersInfoStruct.AllPapersInfo[i].PaperName, globals.AllPapersInfoStruct.AllPapersInfo[i].CompanyName, globals.AllPapersInfoStruct.AllPapersInfo[i].DailyRate, mValueInFloat)

		if err != nil {
			fmt.Print("At writing to db (Exec): ", err.Error())
		}

		defer stmt.Close()

	}
}
