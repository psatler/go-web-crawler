package databaseutils

import (
	"database/sql"
	"fmt"
	"strconv"
	"unicode/utf8"

	//custom package
	"github.com/psatler/go-web-crawler/globals"
)

func parseNonUtf8(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	return s
}

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
		companyName := parseNonUtf8(globals.AllPapersInfoStruct.AllPapersInfo[i].CompanyName)
		_, err = stmt.Exec(globals.AllPapersInfoStruct.AllPapersInfo[i].PaperName, companyName, globals.AllPapersInfoStruct.AllPapersInfo[i].DailyRate, mValueInFloat)

		if err != nil {
			fmt.Print("At writing to db (Exec): ", err.Error())
		}

		defer stmt.Close()

	}
}
