package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

func CSVExport(headers, records []string, filename string) (*os.File, error) {
	substring := strings.Split(filename, ".")
	if strings.Compare(substring[len(substring)-1], "csv") != 0 {
		panic(errors.New("file must be csv"))
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
			return
		}
	}(file)
	writer := csv.NewWriter(file)
	err = writer.Write(headers)
	if err != nil {
		panic(err)
		return nil, err
	}
	err = writer.Write(records)
	if err != nil {
		panic(err)
		return nil, err
	}
	return file, nil
}
