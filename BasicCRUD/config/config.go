package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var DB *gorm.DB
var DB *sqlx.DB
var Cache *redis.Client
var Client *mongo.Client


func ConnectDB() {
	dsn:="user=postgres password=Password123@ database=CarInventory host=localhost port=5432 sslmode=disable"

	db,err := sqlx.Connect("postgres",dsn)
	if err!=nil{
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}

	
	if err := db.Ping(); err != nil {
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
func ConnectMongo(){
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		panic(err)
	}
	Client = client
	fmt.Println("Connected to MongoDB successfully")
}