package mongoimport

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func TestCSVImport(t *testing.T) {
	records_len := len(CSV_Reader("sample.csv"))
	records := CSV_Reader("sample.csv")
	collection := getCollection()
	if collection == nil {
		t.Fatalf("Failed to establish Mongodb collection connection")
	}

	got := CSVImport(collection, records, 1, records_len)
	assert.Equal(t, records_len-1, got, "Expected number of records to be inserted")
}

// test function to get collection
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
