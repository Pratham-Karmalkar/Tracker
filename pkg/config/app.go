package config

//Data base connection is established here

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

//fmt.Println(DB_USER,DB_PASS,DB_PROTOCOL,DB_ADDRESS,DB_PORT,DB_DATABASE)
//"root:mybench@tcp(localhost:3306)/Tracker"
//Data base connector
func Connect() {

	//loads env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading env")
	}

	var (
		DB_USER     = os.Getenv("USER_NAME")
		DB_PASS     = os.Getenv("PASS")
		DB_PROTOCOL = os.Getenv("PROTOCOL")
		DB_ADDRESS  = os.Getenv("ADDRESS")
		DB_PORT     = os.Getenv("PORT")
		DB_DATABASE = os.Getenv("DATABASE")
	)

	d, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_PROTOCOL+"("+DB_ADDRESS+":"+DB_PORT+")"+"/"+DB_DATABASE)
	if err != nil {
		panic(err.Error())

	}
	db = d

}

//Returns Instace of SQL
func GetDB() *sql.DB {
	return db
}
