// package csv
//
// import (
//
//	"github.com/stretchr/testify/assert"
//	"gopkg.in/mgo.v2"
//	"strconv"
//	"testing"
//
// )
//
//	type Sales struct {
//		ItemId   int    `bson:"itemId" json:"itemId"`
//		Name     string `bson:"name" json:"name"`
//		Quantity int    `bson:"quantity" json:"quantity"`
//		Price    int    `bson:"price" json:"price"`
//		Amount   int    `bson:"amount" json:"amount"`
//		Status   string `bson:"status" json:"status"`
//		Date     string `bson:"Date" json:"Date"`
//	}
//
//	func TestCSVExport(t *testing.T) {
//		headers := []string{"ItemId", "Name", "Quantity", "Price", "Amount", "Status", "Date"}
//		records, _ := getDBData()
//		type args struct {
//			Records  []string
//			Headers  []string
//			filename string
//		}
//
//		tests := []struct {
//			name    string
//			args    args
//			want    int
//			wantErr bool
//		}{
//			{
//				name:    "Case 1: Check file size if > 1 byte",
//				args:    args{filename: "tests.csv", Headers: headers, Records: records},
//				want:    len(records),
//				wantErr: false,
//			},
//			{
//				name:    "Case 2: Throw error non csv file passed",
//				args:    args{filename: "tests.json", Headers: headers, Records: records},
//				want:    0,
//				wantErr: true,
//			},
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				got, err := CSVExport(tt.args.Headers, tt.args.filename)
//				if err != nil {
//					t.Fail()
//				}
//				if tt.wantErr {
//					assert.Error(t, err, "Expected an error")
//				} else {
//					assert.NoError(t, err, "Unexpected error")
//				}
//				assert.Equal(t, tt.want, got, "Unexpected value thrown")
//			})
//		}
//	}
//
//	func getDBData() ([]string, error) {
//		session, err := mgo.Dial("localhost")
//		if err != nil {
//			return nil, err
//		}
//		defer session.Close()
//		session.SetMode(mgo.Monotonic, true)
//
//		collection := session.DB("RestaurantDB").C("Sales")
//
//		var results []Sales
//		err = collection.Find(nil).All(&results)
//		if err != nil {
//			return nil, err
//		}
//
//		var records []string
//		for _, result := range results {
//			record := []string{
//				strconv.Itoa(result.ItemId),
//				result.Name,
//				strconv.Itoa(result.Quantity),
//				strconv.Itoa(result.Price),
//				strconv.Itoa(result.Amount),
//				result.Status,
//				result.Date,
//			}
//			records = append(records, record...)
//		}
//
//		return records, nil
//	}
package csv

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
)

func TestExportMongoDataToCSV(t *testing.T) {
	// Connect to the MongoDB database (replace with your MongoDB URI).
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("Failed to connect to MongoDB: %v", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Get a reference to your MongoDB collection.
	collection := client.Database("yourdb").Collection("yourcollection")

	// Query MongoDB to retrieve sample data.
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		t.Errorf("Failed to query MongoDB: %v", err)
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("err closing cursor %v", err)
		}
	}(cursor, context.Background())

	// Convert MongoDB documents to a slice of structs (or maps) as needed.
	var records []interface{}
	for cursor.Next(context.Background()) {
		var doc YourDocument // Replace with your actual document structure
		if err := cursor.Decode(&doc); err != nil {
			t.Errorf("Failed to decode MongoDB document: %v", err)
			return
		}
		records = append(records, doc)
	}

	// Define headers and a filename.
	headers := []string{"Field1", "Field2", "Field3"} // Replace with your actual headers
	filename := "test.csv"

	// Export the retrieved MongoDB data to a CSV file.
	err = ExportToCSV(headers, records, filename)
	if err != nil {
		t.Errorf("ExportToCSV returned an error: %v", err)
	}

	// Now you can read the CSV file and verify its content if needed.
}

func TestConvertToSlice(t *testing.T) {
	// Example data
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Alice", 30}

	// Convert a struct to a slice
	slice, err := convertToSlice(person)
	if err != nil {
		t.Errorf("convertToSlice returned an error: %v", err)
	}

	expectedSlice := []string{"Alice", "30"}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Errorf("Converted slice doesn't match the expected slice: %v", slice)
	}

	// Test an unsupported record type
	unsupported := []int{1, 2, 3}
	_, err = convertToSlice(unsupported)
	if err == nil {
		t.Error("convertToSlice didn't return an error for an unsupported record type")
	}
}
