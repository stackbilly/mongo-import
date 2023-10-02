package mongoimport

import (
	"encoding/csv"
	"log"
	"os"
)

func CSVReader(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening csv file %v", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error reading csv file %v", err)
		return nil, err
	}
	return records, nil
}
