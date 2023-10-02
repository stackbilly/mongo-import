package mongoimport

import (
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			want:    CSV_Reader("sample.csv"), //sample func for testing
			wantErr: false,
		},
		{
			name:    "case2: should return err",
			args:    args{"sample456.csv"},
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

// test function
func CSV_Reader(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return records
}
