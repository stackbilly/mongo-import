# mongo-import go package

This package is meant to be used to convert data from a csv/json file and write it into a mongodb collection

# mongo-import tools
- **csvreader**: Read csv records & entries from a csv file
- **csvimport**: convert data from csv and write into a new mongodb collection
- **jsonreader**: Read json entries from a json file
- **jsonimport**: convert data from json file and write into a new mongodb collection

## overview

-[Install](#installation)
-[Example](#example)
-[License](#license)

## Installation

### Install via go get

Please ensure you have installed Go v1.21 or later

```sh
go get github.com/stackbilly/mongo-import
```
From source code
```sh
git clone https://github.com/stackbilly/mongo-import.git
cd mongo-import
```
## Example
```go
package main

import (
	"context"
	"fmt"
	"github.com/stackbilly/mongo-import/csv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	//read csv records from file
	records, err := csv.CSVReader("csv/sample.csv")
	if err != nil {
		panic(err)
		return
	}
	//connect to mongodb
	serveAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017/?directConnection=true").SetServerAPIOptions(serveAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
		return
	}
	collection := client.Database("admin").Collection("records")
	start := time.Now()
	count := csv.CSVImport(collection, records, 1, len(records))

	fmt.Printf("Inserted %d docs in %v seconds", count, time.Since(start).Seconds())

	/*
				Sample output
		 Inserted 1338 docs in 0.3408692 seconds
	*/
}
```

## License
This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details
