package json

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestJsonImport test function for json mongodb import
func TestJsonImport(t *testing.T) {
	contents, _ := JSONFileReader("test.json")
	type args struct {
		collection *mongo.Collection
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
func getJsonCollection() *mongo.Collection {
	serveAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017/?directConnection=true").SetServerAPIOptions(serveAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
		return nil
	}
	collection := client.Database("admin").Collection("json")
	return collection
}
