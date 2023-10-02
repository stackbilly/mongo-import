package mongoimport

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

func main() {
	records, err := CSVReader("sample.csv")
	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Printf("error establishing connection to mongodb: %v", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	coll := session.DB("test").C("example-test")
	start := time.Now()
	count := CSVImport(coll, records, 1, len(records))
	fmt.Printf("inserted %d records within %v seconds", count, time.Since(start).Seconds())
}
