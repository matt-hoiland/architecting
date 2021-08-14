package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	mongoHost = "localhost"
	mongoPort = "27117" // To not conflict with a host run mongo
)

type HistoryDocument struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Service string             `bson:"service,omitempty" json:"service,omitempty"`
}

func (hd *HistoryDocument) String() string {
	data, err := json.Marshal(hd)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("Hi!")

	uri := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// Prior to this, I stuffed some bland data into my local instance of the db
	// collection arch.history
	archDB := client.Database("arch")
	history := archDB.Collection("history")
	var hd HistoryDocument
	if err = history.FindOne(ctx, bson.M{}).Decode(&hd); err != nil {
		log.Fatal(err)
	}
	log.Println(hd.String())

	log.Info("Successfully connected and pinged.")
}
