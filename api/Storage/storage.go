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

type Storage struct {}


func InitStorage() Storage {
	filePathName, err := getStorageFilepath()
	handleError(err)

	file, err := os.OpenFile(filePathName, os.O_RDWR|os.O_CREATE, 0666)
	handleError(err)

	defer file.Close()

	start, err := file.Stat()
	handleError(err)

	if start.Size() == 0 {
		_, err := file.WriteString("{}")
		handleError(err)
	}

	return Storage{}
}

func (s Storage) GetData() map[string]string {
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

func (s Storage) SaveData(linkList map[string]string) {
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

func handleError(err error) {
	if err != nil {
		panic(fmt.Sprintln("error: ", err.Error()))
	}
}