package Storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const filePathName string = "web/Storage/links.json"

func GetData() map[string]string {
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
