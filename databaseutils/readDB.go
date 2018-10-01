package databaseutils

import (
	"database/sql"
	"log"

	//custom package
	"github.com/psatler/go-web-crawler/globals"
)

type PapersInfo = globals.PapersInfo //kind of a typedef from C

//reads the database and returns a struct with the information stored in the DB
func ReadFromDb() []PapersInfo {
	var mostValuable []PapersInfo //it'll store the values returned from DB

	//vars used in data source are defined in another file in the package
	dataSourceName := username + ":" + password + "@tcp(" + host + ":3306)/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dataSourceName) //open a connection
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT paperName, companyName, dailyRate, marketValue FROM ibovespa")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var info PapersInfo
		// for each row, scan the result into our tag composite object
		err = results.Scan(&info.PaperName, &info.CompanyName, &info.DailyRate, &info.MarketValue)
		if err != nil {
			panic(err.Error())
		}
		// and then print out
		// log.Printf(info.PaperName, info.CompanyName, info.DailyRate, info.MarketValue)
		mostValuable = append(mostValuable, info)
	}

	return mostValuable
}
