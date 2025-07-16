package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB


func ConnectDB() {
	dsn:="user=postgres password=Password123@ database=CarInventory host=localhost port=5432 sslmode=disable"

	db,err:=sql.Open("postgres",dsn)
	if err!=nil{
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Error pinging the database:", err)
		panic(err)
	}
	fmt.Println("Connected to the database successfully")
	DB = db
}