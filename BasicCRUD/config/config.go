package config

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Cache *redis.Client


func ConnectDB() {
	dsn:="user=postgres password=Password123@ database=CarInventory host=localhost port=5432 sslmode=disable"

	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err!=nil{
		fmt.Println("Unable to move to gorm", err)
		panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		fmt.Println("Error pinging the database:", err)
		panic(err)
	}
	fmt.Println("Connected to the database successfully")
	DB = db
}
func ConnectCache() {
	ctx:=context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	cmd:=rdb.Ping(ctx)
	if cmd.Err() != nil {
		fmt.Println("Error connecting to Redis:", cmd.Err())
		panic(cmd.Err())
	}
	fmt.Println("Connected to Redis successfully")
	Cache = rdb
}