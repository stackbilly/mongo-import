package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//
//func CSVExport(headers []string, records []string, filename string) (int, error) {
//	substring := strings.Split(filename, ".")
//	if strings.Compare(substring[len(substring)-1], "csv") != 0 {
//		err := errors.New("file must be csv")
//		fmt.Println(err)
//		return 0, err
//	}
//	file, err := os.Create(filename)
//	if err != nil {
//		panic(err)
//	}
//	defer func(file *os.File) {
//		err := file.Close()
//		if err != nil {
//			fmt.Printf("Error closing file %v", err)
//		}
//	}(file)
//
//	writer := csv.NewWriter(file)
//	defer writer.Flush()
//	if err := writer.Write(headers); err != nil {
//		return 0, err
//	}
//
//	for _, r := range records {
//		err := writer.Write(records)
//		if err != nil {
//			return 0, err
//		}
//	}
//	info, _ := file.Stat()
//	return int(info.Size()), nil
//}
//
//package mongotocsv

// ExportToCSV exports MongoDB documents to a CSV file.
func ExportToCSV(headers []string, records []interface{}, filename string) error {
	// Check if the provided filename has a ".csv" extension.
	if !strings.HasSuffix(filename, ".csv") {
		return errors.New("file must have a .csv extension")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file %s", err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the headers to the CSV file.
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Iterate through the records and write each one to the CSV file.
	for _, record := range records {
		// Convert the record to a []string slice.
		recordSlice, err := convertToSlice(record)
		if err != nil {
			return err
		}

		if err := writer.Write(recordSlice); err != nil {
			return err
		}
	}

	return nil
}

// convertToSlice converts a MongoDB document (struct or map) to a []string slice.
func convertToSlice(record interface{}) ([]string, error) {
	// Use reflection to handle different types of records.
	val := reflect.ValueOf(record)

	// Check if the record is a struct or a map.
	if val.Kind() == reflect.Struct {
		return convertStructToSlice(val)
	} else if val.Kind() == reflect.Map {
		return convertMapToSlice(val)
	}

	return nil, errors.New("unsupported record type")
}

func convertStructToSlice(val reflect.Value) ([]string, error) {
	var recordSlice []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		recordSlice = append(recordSlice, fieldToString(field))
	}
	return recordSlice, nil
}

func convertMapToSlice(val reflect.Value) ([]string, error) {
	var recordSlice []string
	for _, key := range val.MapKeys() {
		value := val.MapIndex(key)
		recordSlice = append(recordSlice, fieldToString(value))
	}
	return recordSlice, nil
}

func fieldToString(field reflect.Value) string {
	switch field.Kind() {
	case reflect.String:
		return field.String()
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return strconv.FormatUint(field.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(field.Float(), 'f', -1, 32)
	default:
		return ""
	}
}
