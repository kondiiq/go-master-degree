package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"os"
	"strconv"
)

var Db *sql.DB //created outside to make it global.

//make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {

	err := godotenv.Load()//by default, it is .env so we don't have to write
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
  //we read our .env file
	host := os.Getenv("POSTGRES_HOST")
  	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")
  // set up postgres sql to open it.
    psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
    host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err) 
	} else {
	Db = db
	fmt.Println("Successfully connected to database!")
	}
}