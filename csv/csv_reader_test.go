package csv

import (
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCSVReader test function -> CSVReader
func TestCSVReader(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name:    "case1: should return csv records as slice",
			args:    args{"sample.csv"},
			want:    CsvReader("sample.csv"), //sample func for testing
			wantErr: false,
		},
		{
			name:    "case2: should return err for non-existent file",
			args:    args{"sample456.csv"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case3: should return error for non csv file passed",
			args:    args{filename: "test.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CSVReader(tt.args.filename)
			if tt.wantErr {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
			assert.Equal(t, tt.want, got, "CSVReader() result not expected")
		})
	}
}

// sample test function for CSVReader
func CsvReader(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return records
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCSVReader$ github.com/Livingstone-Billy/mongo-import

// === RUN   TestCSVReader
// === RUN   TestCSVReader/case1:_should_return_csv_records_as_slice
// --- PASS: TestCSVReader/case1:_should_return_csv_records_as_slice (0.00s)
// === RUN   TestCSVReader/case2:_should_return_err_for_non-existent_file
// 2023/10/10 20:13:42 Error opening csv file open sample456.csv: The system cannot find the file specified.
// --- PASS: TestCSVReader/case2:_should_return_err_for_non-existent_file (0.00s)
// === RUN   TestCSVReader/case3:_should_return_error_for_non_csv_file_passed
// --- PASS: TestCSVReader/case3:_should_return_error_for_non_csv_file_passed (0.00s)
// --- PASS: TestCSVReader (0.01s)
// PASS
// ok      github.com/Livingstone-Billy/mongo-import       0.388s
