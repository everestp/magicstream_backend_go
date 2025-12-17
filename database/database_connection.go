package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
 func DBInstance()*mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to Find  .env files")

	}
	MongoDB := os.Getenv("MONGODB_URI")
	if MongoDB == ""{
		log.Fatal("MONGODB_URI not set!")

	}

	fmt.Println("Mongodb URI",MongoDB)
	clientOptions := options.Client().ApplyURI(MongoDB)
	client , err := mongo.Connect(clientOptions)
	if err != nil {  
		return nil
	}
return  client

 }

   var Client *mongo.Client = DBInstance()

 func OpenCollection(collectionName  string) *mongo.Collection {
	 err := godotenv.Load(".env")
	 if err != nil {
		log.Println("Warning : unable to  fund .env file")

	 }
	 databaseName := os.Getenv("DATABASE_NAME")
	fmt.Println("DATABASE_NAME:",databaseName)

	// opeining the collection
	collection := Client.Database(databaseName).Collection(collectionName)
	if  collection == nil {
		return nil

	}
return  collection
}