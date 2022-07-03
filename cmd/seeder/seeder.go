package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"grabjobs/domain"
	"grabjobs/internal/constants"
	"grabjobs/internal/helpers"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	godotenv.Load()
}

// SeedData is used to populate data into the database from the csv files
func SeedData() {
	csvLines, err := helpers.LoadCSV("location_data_2000.csv")
	if err != nil {
		log.Fatal(err)
	}
	var dataToSeed []interface{}
	err = DecodeCSV(csvLines, &dataToSeed)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	//Verify MongoDB connection
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	database := client.Database(os.Getenv("MONGO_DATABASE"))
	database.Collection(constants.JobsCollection).Drop(context.TODO())
	opts := options.InsertMany().SetOrdered(false)
	res, err := database.Collection(constants.JobsCollection).InsertMany(context.TODO(), dataToSeed, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted  %v documents\n", len(res.InsertedIDs))
	coll := database.Collection(constants.JobsCollection)
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Cfreation of index for faster queries and from mongo documentation it is preferrable to use 2dsphere
	mod := mongo.IndexModel{
		Keys:    bson.M{"location": "2dsphere"},
		Options: nil,
	}
	ind, err := coll.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("CreateOne() index:", ind)
	}
}

// DecodeCSV is used to decode the csvlines into an array of Jobs struct
func DecodeCSV(csvLines [][]string, dataToDecode *[]interface{}) error {
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		longitude, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return err
		}
		latitude, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return err
		}
		coordinates := make([]float64, 2)
		coordinates[0] = longitude
		coordinates[1] = latitude
		emp := domain.Jobs{
			Title: line[0],
			Location: domain.GeoJSON{
				Type:        "Point",
				Coordinates: coordinates,
			},
		}
		*dataToDecode = append(*dataToDecode, emp)
	}
	return nil
}
