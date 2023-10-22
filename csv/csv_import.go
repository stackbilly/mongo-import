package csv

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CSVImport func to import csv records to mongodb
func CSVImport(collection *mgo.Collection, records [][]string, start, end int) int {
	sliceOfRecords := make([][]string, len(records))
	copy(sliceOfRecords, records)
	colNames := sliceOfRecords[0]

	bulk := collection.Bulk()
	bulk.Unordered()

	for i := start; i < end; i++ {
		bsonData := make(bson.M)
		for j := 0; j < len(colNames); j++ {
			bsonData[colNames[j]] = records[i][j]
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
