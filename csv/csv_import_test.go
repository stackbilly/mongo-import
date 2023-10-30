package csv

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCSVImport test func for CSVImport
func TestCSVImport(t *testing.T) {
	records := CsvReader("sample.csv")
	recordsLen := len(records)
	collection, _ := getCollection()
	if collection == nil {
		t.Fatalf("Failed to establish Mongodb collection connection")
	}

	got := CSVImport(collection, records, 1, recordsLen)
	assert.Equal(t, recordsLen-1, got, "Expected number of records to be inserted")
}

// test function to get mongo collection
func getCollection() (*mongo.Collection, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017/?directConnection=true").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
		return nil, err
	}
	collection := client.Database("admin").Collection("tests")
	return collection, nil
}
