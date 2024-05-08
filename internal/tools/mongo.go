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
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://madmin:password@gserver:27017/"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := mongoClient.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
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
			mongo.NewReplaceOneModel().SetFilter(bson.D{{Key: "_id", Value: game.ID}}).SetUpsert(true).SetReplacement(game))
	}

	opts := options.BulkWrite().SetOrdered(true)

	results, err := coll.BulkWrite(context.TODO(), models, opts)

	if err != nil {
		panic(err)
 	}

	fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
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
		output, err := json.MarshalIndent(event["fullDocument"], "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	if err := cs.Err(); err != nil {
		panic(err)
	}
}