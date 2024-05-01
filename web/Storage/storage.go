package Storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func GetData() map[string]string {
	filePathName, err := getStorageFilepath()
	if err != nil {
		panic(fmt.Sprintln("error on get filePath", err.Error()))
	}

	jsonFile, err := os.Open(filePathName)
	if err != nil {
		fmt.Println("error on open json: ", err.Error())
	}

	defer jsonFile.Close()

	mapValues, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(fmt.Sprintln("error on converting json to map: ", err.Error()))
	}

	linkList := make(map[string]string)

	if err := json.Unmarshal(mapValues, &linkList); err != nil {
		panic(fmt.Sprintln("error on open json: ", err.Error()))
	}
	return linkList
}

func SaveData(linkList map[string]string) {
	filePathName, err := getStorageFilepath()
	if err != nil {
		panic(fmt.Sprintln("error on get filePath", err.Error()))
	}

	file, err := os.Create(filePathName)

	if err != nil {
		panic("error on creatin file")
	}

	convertedMap, err := json.Marshal(linkList)

	if err != nil {
		panic("error to convert json")
	}

	defer file.Close()

	_, err = file.Write(convertedMap)

	if err != nil {
		panic("error on writing on file")
	}
}

func getStorageFilepath() (string, error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return "", errors.New("unable to get the filename")
	}

	return fmt.Sprintf("%v/links.json",filepath.Dir(filename)), nil
}
