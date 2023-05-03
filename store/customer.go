package store

import (
	"context"
	"os"
	"time"
    "log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
    conn *mongo.Client
    coll *mongo.Collection 
}

type Result struct {
    ID primitive.ObjectID `json:"id" bson:"_id"`
    Username string `json:"username"`
    Name string `json:"name"`
    Address string `json:"address"`
    Birthdate time.Time `json:"birthdate" bson:"birthdate"`
    Email string `json:"email"`
    Accounts []int32 `json:"accounts"`
    Tier_and_details bson.D `json:"tier_and_details"`
}

func NewDB() (*Database,error) {
    var err error;
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    uri := os.Getenv("mongo_uri")
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
    db := &Database{}
    db.conn , err = mongo.Connect(context.TODO(),opts) 
    if err != nil {
        return db , err 
    }
    db.coll = db.conn.Database("sample_analytics").Collection("customers")
    return db,nil
}

func (db *Database) Retrievebyname(name string) (Result,error) {
    var fltr interface{}
    filter_error := bson.UnmarshalExtJSON([]byte(name),true,&fltr)
    if filter_error != nil {
        log.Println(filter_error)
    }
    filter := bson.D{{"name",fltr}}
    var result Result
    err := db.coll.FindOne(context.TODO(),filter).Decode(&result)
    if err != nil {
        return result, err 
    }
    return result,nil
}








