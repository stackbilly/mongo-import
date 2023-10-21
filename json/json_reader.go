package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// JSONFileReader function json file reader
func JSONFileReader(filename string) (map[string]interface{}, error) {
	substring := strings.Split(filename, ".")
	var extension string
	if len(substring) > 1 {
		extension = substring[len(substring)-1]
	} else {
		return nil, errors.New("invalid file passed")
	}
	if strings.Compare("json", extension) == 0 {
		jsonfile, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("err opening file \n%v", err)
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
			return nil, fmt.Errorf("error reading file stats \n%v", err)
		}
		fileSize := fileInfo.Size()

		byteValue := make([]byte, fileSize)
		bytesRead, err := jsonfile.Read(byteValue)
		if err != nil {
			log.Printf("error reading %d bytes\n%s", bytesRead, err)
			return nil, err
		}
		var contents map[string]interface{}
		err = json.Unmarshal(byteValue, &contents)
		if err != nil {
			return nil, fmt.Errorf("err parsing json from file \n%v", err)
		}

		return contents, nil
	} else {
		return nil, errors.New("json file required")
	}
}
