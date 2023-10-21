package mongoimport

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

// TestJsonImport test function for json mongodb import
func TestJsonImport(t *testing.T) {
	contents, _ := JSONFileReader("test.json")
	type args struct {
		collection *mgo.Collection
		contents   map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return collection count after insertion",
			args: args{collection: getJsonCollection(), contents: contents},
			want: len(contents),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := JSONImport(tt.args.collection, tt.args.contents)
			assert.Equal(t, tt.want, got)
		})
	}
}

// test function to get mongo collection
func getJsonCollection() *mgo.Collection {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Printf("Failed to establish connection to MongoDB: %v", err)
		return nil
	}
	defer session.SetMode(mgo.Monotonic, true)
	coll := session.DB("jsontest").C("json-colls")
	return coll
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestJsonImport$ github.com/Livingstone-Billy/mongo-import

// === RUN   TestJsonImport
// === RUN   TestJsonImport/should_return_collection_count_after_insertion
// --- PASS: TestJsonImport/should_return_collection_count_after_insertion (0.40s)
// --- PASS: TestJsonImport (0.43s)
// PASS
// ok      github.com/Livingstone-Billy/mongo-import       0.499s
