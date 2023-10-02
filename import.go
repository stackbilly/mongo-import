package mongoimport

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CSVImport(collection *mgo.Collection, records [][]string, start, end int) int {
	slice_of_records := make([][]string, len(records))
	copy(slice_of_records, records)
	col_names := slice_of_records[0]

	bulk := collection.Bulk()
	bulk.Unordered()

	for i := start; i < end; i++ {
		bsonData := make(bson.M)
		for j := 0; j < len(col_names); j++ {
			bsonData[col_names[j]] = records[i][j]
		}
		bulk.Insert(bsonData)
	}
	_, err := bulk.Run()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	count, _ := collection.Count()
	return count
}

//**********json import implementation coming soon***********
