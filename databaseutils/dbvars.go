package databaseutils

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// +-------------+--------------+------+-----+---------+-------+
// | Field       | Type         | Null | Key | Default | Extra |
// +-------------+--------------+------+-----+---------+-------+
// | paperName   | varchar(255) | YES  |     | NULL    |       |
// | companyName | varchar(255) | YES  |     | NULL    |       |
// | dailyRate   | varchar(20)  | YES  |     | NULL    |       |
// | marketValue | float        | YES  |     | NULL    |       |
// +-------------+--------------+------+-----+---------+-------+

var username = envVar("DB_USERNAME", "root")
var password = envVar("DB_PASSWORD", "root")
var host = envVar("DB_HOST", "localhost")
var dbname = envVar("DB_NAME", "demodb")
var port = envVar("DB_PORT", "3306")

func envVar(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
