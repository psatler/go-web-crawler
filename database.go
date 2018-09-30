package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// +-------------+--------------+------+-----+---------+-------+
// | Field       | Type         | Null | Key | Default | Extra |
// +-------------+--------------+------+-----+---------+-------+
// | paperName   | varchar(255) | YES  |     | NULL    |       |
// | companyName | varchar(255) | YES  |     | NULL    |       |
// | dailyRate   | varchar(20)  | YES  |     | NULL    |       |
// | marketValue | float        | YES  |     | NULL    |       |
// +-------------+--------------+------+-----+---------+-------+

func envVar(envVar string, defaultVar string) string {
	foo, ok := os.LookupEnv(envVar)
	if ok {
		return foo
	} else {
		return defaultVar
	}
}

var username = envVar("DB_USERNAME", "root")
var password = envVar("DB_PASSWORD", "pablo2908")
var host = envVar("DB_HOST", "localhost")
var dbname = envVar("DB_NAME", "demodb")
var port = envVar("DB_PORT", "3306")

//writes the result to a mysql database
func WriteToDb() {
	//db requests
	dataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dataSourceName) //open a connection
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	for i := 0; i < 10; i++ {
		mValueInFloat := strconv.FormatFloat(allPapersInfo[i].marketValue, 'f', 0, 64)
		// fmt.Printf("\n#%d - \t Company: %s \n \t Market Value: %s \n", i, allPapersInfo[i].companyName, mValueInFloat)

		stmt, err := db.Prepare("insert into ibovespa (paperName, companyName, dailyRate, marketValue) values(?,?,?,?);")
		if err != nil {
			fmt.Print("At writing to db (Prepare): ", err.Error())
		}
		_, err = stmt.Exec(allPapersInfo[i].paperName, allPapersInfo[i].companyName, allPapersInfo[i].dailyRate, mValueInFloat)

		if err != nil {
			fmt.Print("At writing to db (Exec): ", err.Error())
		}

		defer stmt.Close()

	}
}

//reads the database and returns a struct with the information stored in the DB
func ReadFromDb() []PapersInfo {
	var mostValuable []PapersInfo //it'll store the values returned from DB

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
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var info PapersInfo
		// for each row, scan the result into our tag composite object
		err = results.Scan(&info.paperName, &info.companyName, &info.dailyRate, &info.marketValue)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(info.paperName, info.companyName, info.dailyRate, info.marketValue)
		mostValuable = append(mostValuable, info)
	}

	return mostValuable
}
