package mongoimport

import (
	"log"

	"gopkg.in/mgo.v2"
)

// JSONImport function import/write json contents into a mongodb collection
func JSONImport(collection *mgo.Collection, contents map[string]interface{}) (int, error) {
	bulk := collection.Bulk()
	bulk.Unordered()

	bulk.Insert(contents)
	_, err := bulk.Run()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	count, err := collection.Count()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return count, nil
}
