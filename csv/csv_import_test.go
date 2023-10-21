package csv

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

// TestCSVImport test func for CSVImport
func TestCSVImport(t *testing.T) {
	records := CsvReader("sample.csv")
	recordsLen := len(records)
	collection := getCollection()
	if collection == nil {
		t.Fatalf("Failed to establish Mongodb collection connection")
	}

	got := CSVImport(collection, records, 1, recordsLen)
	assert.Equal(t, recordsLen-1, got, "Expected number of records to be inserted")
}

// test function to get mongo collection
func getCollection() *mgo.Collection {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Printf("Failed to establish connection to MongoDB: %v", err)
		return nil
	}
	defer session.SetMode(mgo.Monotonic, true)
	coll := session.DB("test").C("test-col")
	return coll
}