package main

import (
	"context"
	"fmt"
	"github.com/stackbilly/mongo-import/csv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	//read csv records from file
	records, err := csv.CSVReader("csv/sample.csv")
	if err != nil {
		panic(err)
		return
	}
	//connect to mongodb
	serveAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017/?directConnection=true").SetServerAPIOptions(serveAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
		return
	}
	collection := client.Database("admin").Collection("records")
	start := time.Now()
	count := csv.CSVImport(collection, records, 1, len(records))

	fmt.Printf("Inserted %d docs in %v seconds", count, time.Since(start).Seconds())

	/*
				Sample output
		 Inserted 1338 docs in 0.3408692 seconds
	*/
}
