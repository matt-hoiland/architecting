package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/matt-hoiland/architecting/data"
	"github.com/matt-hoiland/architecting/internal/api/auth"
	"github.com/matt-hoiland/architecting/lib/flag"
	"github.com/matt-hoiland/architecting/lib/logging"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ServiceName         = "AuthService"
	DefaultLogLevel     = "INFO"
	DefaultMongoHost    = "localhost"
	DefaultMongoPort    = "27117" // NOTE: See launch-mongo.sh for why it is not 27017
	DefaultMongoTimeout = 5       // seconds
)

var (
	logLevel = flag.String("log-level", "LOG_LEVEL", DefaultLogLevel, "logging level for service")
	mhost    = flag.String("mongo-host", "MONGO_HOST", DefaultMongoHost, "hostname of mongodb server")
	mport    = flag.String("mongo-port", "MONGO_PORT", DefaultMongoPort, "port of mongodb server")
	mtimeout = flag.Int("mongo-timeout", "MONGO_TIMEOUT", DefaultMongoTimeout, "number of seconds to use for mongo's ServerSelectionTimeout value")
)

func main() {
	flag.Parse()
	log.SetLevel(logging.LevelFromString(*logLevel))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info(ServiceName + " starting ...")
	defer log.Info(ServiceName + " closing ...")

	mongoURI := fmt.Sprintf("mongodb://%s:%s", *mhost, *mport)
	opts := options.Client().
		ApplyURI(mongoURI).
		SetAppName(ServiceName).
		SetServerSelectionTimeout(time.Duration(*mtimeout) * time.Second)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"mongoURI":   mongoURI,
			"db":         data.AuthDatabaseName,
			"collection": data.CredentialsCollectionName,
		}).Fatalf("Error connecting to mongodb")
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()

	db := client.Database(data.AuthDatabaseName)
	specs, err := db.ListCollectionSpecifications(ctx, primitive.M{})
	if err != nil {
		log.Error(err)
	}
	spec := specs[0]
	var options primitive.M
	err = bson.Unmarshal(spec.Options, &options)
	if err != nil {
		log.Error(err)
	}

	// err = db.CreateCollection(ctx, auth.CredentialsCollection, options.CreateCollection().SetValidator(data.AuthCredentialsSchema))
	// if err != nil {
	// 	log.Error(err)
	// }

	collection := client.Database(data.AuthDatabaseName).Collection(data.CredentialsCollectionName)
	authAPI := auth.NewAuthAPI(collection)
	debug(ctx, authAPI)

	http.HandleFunc("/health/mongodb", makeMongoHealthCheckHandler(ctx, client))

	// if err = http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Panic(err)
	// }
}

func makeMongoHealthCheckHandler(ctx context.Context, client *mongo.Client) http.HandlerFunc {
	type response struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
		Error   string `json:"error,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := response{
			Code:    http.StatusOK,
			Message: "OK on mongodb",
		}

		err := client.Ping(ctx, readpref.Primary())
		if err != nil {
			res.Code = http.StatusServiceUnavailable
			res.Message = "mongodb unreachable"
			res.Error = err.Error()
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("error encoding response")
		}
	}
}
