package json

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func JSONImport(collection *mongo.Collection, contents map[string]interface{}) (int, error) {
	var bulkModels []mongo.WriteModel

	for i := 0; i < len(contents); i++ {
		bulkModels = append(bulkModels, mongo.NewInsertOneModel().SetDocument(contents))
	}

	ctx := context.TODO()
	_, err := collection.BulkWrite(ctx, bulkModels)

	if err != nil {
		panic(err)
		return 0, err
	}
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		panic(err)
		return 0, err
	}
	return int(count), nil
}
