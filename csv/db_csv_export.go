package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

func CSVExport(headers []string, records []string, filename string) (int, error) {
	substring := strings.Split(filename, ".")
	if strings.Compare(substring[len(substring)-1], "csv") != 0 {
		err := errors.New("file must be csv")
		//panic()
		return 0, err
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		return 0, err
	}
	if err := writer.Write(records); err != nil {
		return 0, err
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return 0, err
	}
	info, _ := file.Stat()
	return int(info.Size()), nil
}

////ExtHeaders func  extracts headers from field names assume all docs are identical
//func ExtHeaders()[]string{
//
//}
