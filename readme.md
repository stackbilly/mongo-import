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
	"fmt"
	"log"
	"time"

	mongoimport "github.com/stackbilly/mongo-import" //import package
	"gopkg.in/mgo.v2"
)

func main() {
	records, err := mongoimport.CSVReader("sample.csv") //read csv file records
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	session, err := mgo.Dial("localhost") //start mongodb session
	if err != nil {
		log.Printf("mongo err: %v", err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	coll := session.DB("test").C("example-coll")
	start := time.Now()
    //write records to mongodb collection
	count := mongoimport.CSVImport(coll, records, 1, len(records)) 

	fmt.Printf("Inserted %d docs in %v seconds", count, time.Since(start).Seconds())

	// Inserted 1338 docs in 0.3729399 seconds
}
```

## License
This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details
