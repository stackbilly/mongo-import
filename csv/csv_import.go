package csv

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CSVImport func to import csv records to mongodb

func CSVImport(collection *mongo.Collection, records [][]string, start, end int) int {
	sliceOfRecords := make([][]string, len(records))
	copy(sliceOfRecords, records)

	colNames := sliceOfRecords[0]

	var bulkModels []mongo.WriteModel

	for i := start; i < end; i++ {
		bsonData := make(bson.M)
		for j := 0; j < len(colNames); j++ {
			bsonData[colNames[j]] = records[i][j]
		}
		bulkModels = append(bulkModels, mongo.NewInsertOneModel().SetDocument(bsonData))
	}

	ctx := context.TODO()

	_, err := collection.BulkWrite(ctx, bulkModels)

	if err != nil {
		panic(err)
		return 0
	}

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		panic(err)
		return 0
	}
	return int(count)
}
