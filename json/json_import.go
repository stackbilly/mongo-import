package json

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// JSONImport function import/write json contents into a mongodb collection
//func JSONImport(collection *mgo.Collection, contents map[string]interface{}) (int, error) {
//	bulk := collection.Bulk()
//	bulk.Unordered()
//
//	bulk.Insert(contents)
//	_, err := bulk.Run()
//	if err != nil {
//		log.Fatal(err)
//		return 0, err
//	}
//	count, err := collection.Count()
//	if err != nil {
//		log.Fatal(err)
//		return 0, err
//	}
//	return count, nil
//}

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
