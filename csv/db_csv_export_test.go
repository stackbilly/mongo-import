package csv

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"testing"
)

type Sales struct {
	Id       bson.ObjectId `bson:"_id omitempty"`
	ItemId   int           `json:"itemId"`
	Name     string        `json:"name"`
	Quantity int           `json:"quantity"`
	Price    int           `json:"price"`
	Amount   int           `json:"amount"`
	Status   string        `json:"status"`
	Date     string        `json:"date"`
}

func TestCSVExport(t *testing.T) {
	headers := []string{"ItemId", "Name", "Quantity", "Price", "Amount", "Status", "Date"}
	records, _ := getDBData()
	type args struct {
		Records  []string
		Headers  []string
		filename string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Case 1: Check file size if > 1 byte",
			args:    args{filename: "tests.csv", Headers: headers, Records: records},
			want:    len(records),
			wantErr: false,
		},
		{
			name:    "Case 2: Throw error non csv file passed",
			args:    args{filename: "tests.json", Headers: headers, Records: records},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CSVExport(tt.args.Headers, tt.args.Records, tt.args.filename)
			if err != nil {
				t.Fail()
			}
			if tt.wantErr {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
			assert.Equal(t, tt.want, got, "Unexpected value thrown")
		})
	}
}

func getDBData() ([]string, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("RestaurantDB").C("Sales")

	var results []Sales
	err = collection.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, result := range results {
		record := []string{
			result.Id.Hex(), // Assuming ID is a bson.ObjectId
			strconv.Itoa(result.ItemId),
			result.Name,
			strconv.Itoa(result.Quantity),
			strconv.Itoa(result.Price),
			strconv.Itoa(result.Amount),
			result.Status,
			result.Date,
		}
		records = append(records, record...)
	}

	return records, nil
}
