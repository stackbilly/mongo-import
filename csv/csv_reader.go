package csv

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
)

// CSVReader func for csv file reader
func CSVReader(filename string) ([][]string, error) {
	substring := strings.Split(filename, ".")
	var extension string
	if len(substring) > 1 {
		extension = substring[len(substring)-1]
	} else {
		return nil, errors.New("invalid file passed")
	}
	if strings.Compare("csv", extension) == 0 {
		file, err := os.Open(filename)
		if err != nil {
			log.Printf("Error opening csv file %v", err)
			return nil, err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Printf("%s", err)
				return
			}
		}(file)

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			log.Printf("Error reading csv file %v", err)
			return nil, err
		}
		return records, nil
	} else {
		return nil, errors.New("csv file required")
	}
}
