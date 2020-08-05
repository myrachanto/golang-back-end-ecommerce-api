package repository

import(
	"os"
	"log"
	"fmt"
	"context"
	"github.com/joho/godotenv"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/myrachanto/asokomonolith/httperrors"
	// "github.com/myrachanto/asokomonolith/support"
)
var ctx = context.TODO()

func Mongoclient()(*mongo.Client, *httperrors.HttpError){
	err7 := godotenv.Load()
	if err7 != nil {
		return nil, httperrors.NewBadRequestError("error loading env file")
	}
	host := os.Getenv("Mongohost")
	clientOptions := options.Client().ApplyURI(host)
	client, err6 := mongo.Connect(ctx, clientOptions)
		if err6 != nil {
		return nil, httperrors.NewBadRequestError("Could not connect to mongodb")
	}
	err8 := client.Ping(ctx, nil)
	if err8 != nil {
		log.Fatal(err8)
	}
	return client, nil
}
func Mongodb()(*mongo.Database, *httperrors.HttpError){
	err7 := godotenv.Load()
	if err7 != nil {
		return nil, httperrors.NewBadRequestError("error loading env file")
	}
	mongodb := os.Getenv("MongodbName")
	client, err1 := Mongoclient()
	db := client.Database(mongodb)
	return db, err1
}
func DbClose(client *mongo.Client){
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
// func SearchQuery(Ser *support.Search,  val []interface{})(values []interface{}, err *httperrors.HttpError){
// 	c, t := Mongoclient();if t != nil {
// 		return nil, t
// 	}
// 	db, e := Mongodb();if e != nil {
// 		return nil, e
// 	}
// 	collection := db.Collection("user")
// 	filter := Ser.Column, 
// 		bson.D{{
// 			Ser.Search_operator, 
// 			bson.A{"Alice", "Bob"}
// 		}}
// 	cur, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
// 	}
// 	if err != nil { 
// 		return nil,	httperrors.NewNotFoundError("no results found")
// 	}
// 	defer cur.Close(ctx)
// 	for cur.Next(ctx) {
// 	err := cur.Decode(&users)
// 		if err != nil { 
// 			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
// 		}
// 	// do something with result....
// 	}
// 	if err := cur.Err(); err != nil {
// 		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
// 	}	
// 	DbClose(c)
//     return users, nil

// }