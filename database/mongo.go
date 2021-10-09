package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "YOUR_URI_HERE"

func Connect() (client *mongo.Client, ctx context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func CreateUser(value interface{}) {
	client, ctx := Connect()
	UserDB := client.Database("crudApi").Collection("users")
	_, err := UserDB.InsertOne(ctx, value)
	if err != nil {
		return
	}
	fmt.Println("created")
}

func GetUserFromID(id int) {
	client, ctx := Connect()
	UserDB := client.Database("crudApi").Collection("users")
	filter := UserDB.FindOne(ctx, bson.M{"id": id})
	fmt.Println(filter)
}

func CreatePost(value interface{}) {
	client, ctx := Connect()
	PostDB := client.Database("crudApi").Collection("posts")
	_, err := PostDB.InsertOne(ctx, value)
	if err != nil {
		return
	}
	fmt.Println("created")
}

func GetPostFromID(id int) {
	client, ctx := Connect()
	UserDB := client.Database("crudApi").Collection("users")
	filter := UserDB.FindOne(ctx, bson.M{"id": id})
	fmt.Println(filter)
}

func GetAllPostsOfUser(id int) {
	client, ctx := Connect()
	UserDB := client.Database("crudApi").Collection("users")
	filter, err := UserDB.Find(ctx, bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	var allPosts []bson.M
	if err = filter.All(ctx, &allPosts);err != nil {
		log.Fatal(err)
	}
	fmt.Println(allPosts)
}