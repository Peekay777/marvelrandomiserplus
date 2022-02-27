package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadData(fileName string) (*Data, error) {
	dataFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("file not found %v", fileName)
	}
	defer dataFile.Close()

	byteValue, _ := ioutil.ReadAll(dataFile)
	var data Data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json file %v", fileName)
	}
	return &data, nil
}
