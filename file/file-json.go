package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func CreateJsonBasedJson(filePath string) {
	fmt.Printf("Creating json file based on exising Yml config\n")
	filePathDest := filepath.Join("", filePath)
	filePathSrc := filepath.Join("", "config-sample.yml")
	file, err := os.Create(filePathDest)
	defer file.Close()
	ymlConfigFile, err := os.ReadFile(filePathSrc)
	CheckError(err)
	var yamlStruct Configuration
	err = yaml.Unmarshal(ymlConfigFile, &yamlStruct)
	CheckError(err)
	rawJsonData, err := json.MarshalIndent(yamlStruct, "", "\t")
	CheckError(err)
	file.Write(rawJsonData)
	CheckError(err)
	fmt.Println("Json data written")
}
