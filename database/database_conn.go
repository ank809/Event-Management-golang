package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBInstance()

func DBInstance() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error in Dot env")
		fmt.Println(err)
	}
	url := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Println("Error hi")
		fmt.Println(err.Error())
	}
	return client
}

func OpenCollection(conn *mongo.Client, coll string) *mongo.Collection {
	collection := conn.Database("Users").Collection(coll)
	return collection
}
