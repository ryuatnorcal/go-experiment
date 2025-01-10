package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := connection()
	getData(client, err)
	insertData(client, err)
	updateData(client, err)
	deleteData(client, err)
	endConnection(client)

}
func endConnection(client *mongo.Client) {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
func connection() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client, err
}
func getData(client *mongo.Client, err error) {

	coll := client.Database("go-mongo").Collection("movies")
	title := "Breaking Bad"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"Title", title}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
func insertData(client *mongo.Client, err error) {
	coll := client.Database("go-mongo").Collection("movies")
	result, err := coll.InsertOne(context.TODO(), bson.D{
		{"Title", "The Matrix"},
		{"Year", "1999"},
		{"Runtime", "136 min"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}

func updateData(client *mongo.Client, err error) {
	coll := client.Database("go-mongo").Collection("movies")
	result, err := coll.UpdateOne(
		context.TODO(),
		bson.D{{"Title", "The Matrix"}},
		bson.D{
			{"$set", bson.D{{"Year", "xxxx"}}},
		})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n %v", result.MatchedCount, result.ModifiedCount)
}

func deleteData(client *mongo.Client, err error) {
	coll := client.Database("go-mongo").Collection("movies")
	result, err := coll.DeleteOne(
		context.TODO(),
		bson.D{{"Title", "The Matrix"}},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted %v document(s)\n", result.DeletedCount)
}
