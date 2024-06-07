package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ghostcrab/mlb-dc-server/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DoMongo() {
	fmt.Println("GetGames Begin")

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://madmin:password@gserver:27017/"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := mongoClient.Database("mlb_dc").Collection("games")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No documents found\n")
		return
	}

	if err != nil {
		panic(err)
	}

	var results []types.MLBGame

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println("GetGames Called")

	jsonData, err := json.Marshal(results[0:5])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}

func AddGames(games []types.MLBGame) {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://madmin:password@gserver:27017/"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := mongoClient.Database("mlb_dc").Collection("games")

	var models []mongo.WriteModel

	for _, game := range games {
		models = append(models, 
			mongo.NewReplaceOneModel().SetFilter(bson.D{{Key: "_id", Value: game.UID}}).SetUpsert(true).SetReplacement(game))
	}

	opts := options.BulkWrite().SetOrdered(true)

	results, err := coll.BulkWrite(context.TODO(), models, opts)

	if err != nil {
		panic(err)
 	}

	fmt.Printf("%d Games; %d Updated; %d Inserted;\n", len(games), results.ModifiedCount, results.UpsertedCount)
}

func WatchForChanges() {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://madmin:password@gserver:27017/"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := mongoClient.Database("mlb_dc").Collection("games")

	var pipeline mongo.Pipeline

	// Creates a change stream that receives change events
	cs, err := coll.Watch(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}
	defer cs.Close(context.TODO())

	fmt.Println("Waiting For Change Events. Insert something in MongoDB!")

	// Prints a message each time the change stream receives an event
	for cs.Next(context.TODO()) {
		var event bson.M
		if err := cs.Decode(&event); err != nil {
			panic(err)
		}

		_, err := json.MarshalIndent(event["fullDocument"], "", "    ")
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%s\n", output)
	}
	if err := cs.Err(); err != nil {
		panic(err)
	}
}