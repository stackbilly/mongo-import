package mongoimport

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestJsonReader test function for json file reader
func TestJsonReader(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "case 1 should return json entries as a map",
			args:    args{filename: "test.json"},
			want:    JSONFile_Reader("test.json"),
			wantErr: false,
		},
		{
			name:    "case 2 throw error for non json file passed",
			args:    args{filename: "sample.csv"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case 3 throw error for non-existent file",
			args:    args{filename: "nonexistent.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := JSONFileReader(tt.args.filename)
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
			assert.Equal(t, tt.want, got, "JSONFileReader() result not expected")
		})
	}
}

// JSONFile_Reader test function output
func JSONFile_Reader(filename string) map[string]interface{} {
	jsonfile, err := os.Open(filename)
	if err != nil {
		return nil
	}
	//closure
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Printf("err: \n%s", err)
			return
		}
	}(jsonfile)

	fileInfo, err := jsonfile.Stat()
	if err != nil {
		return nil
	}
	fileSize := fileInfo.Size()

	byteValue := make([]byte, fileSize)
	_, err = jsonfile.Read(byteValue)
	if err != nil {
		return nil
	}
	var contents map[string]interface{}
	err = json.Unmarshal(byteValue, &contents)
	if err != nil {
		return nil
	}

	return contents
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestJsonReader$ github.com/Livingstone-Billy/mongo-import

// === RUN   TestJsonReader
// === RUN   TestJsonReader/case_1_should_return_json_entries_as_a_map
// --- PASS: TestJsonReader/case_1_should_return_json_entries_as_a_map (0.00s)
// === RUN   TestJsonReader/case_2_throw_error_for_non_json_file_passed
// --- PASS: TestJsonReader/case_2_throw_error_for_non_json_file_passed (0.00s)
// === RUN   TestJsonReader/case_3_throw_error_for_non-existent_file
// --- PASS: TestJsonReader/case_3_throw_error_for_non-existent_file (0.00s)
// --- PASS: TestJsonReader (0.00s)
// PASS
// ok      github.com/Livingstone-Billy/mongo-import       0.416s
