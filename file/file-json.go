package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func UnmarshalYAMLFile(filePath string) (map[interface{}]interface{}, error) {
	ymlDataInterface := make(map[interface{}]interface{})
	file, err := os.Open(filePath)
	CheckError(err)
	defer file.Close()
	stat, err := file.Stat()
	CheckError(err)
	bs := make([]byte, stat.Size())
	bufio.NewReader(file).Read(bs)
	CheckError(err)
	err = yaml.Unmarshal(bs, &ymlDataInterface)
	return ymlDataInterface, nil
}

func cleanYml(yamlData map[interface{}]interface{}) map[string]interface{} {
	cleanYmlMapping := make(map[string]interface{})
	for key, value := range yamlData {
		assertedKey := key.(string)
		cleanYmlMapping[assertedKey] = value
		assertedMapVal, isInterfaceMapType := value.(map[interface{}]interface{})
		assertedSliceVal, isInterfaceSliceType := value.([]interface{})
		if isInterfaceMapType {
			cleanInnerMap := cleanYml(assertedMapVal)
			cleanYmlMapping[assertedKey] = cleanInnerMap
		}
		if isInterfaceSliceType {
			for _, item := range assertedSliceVal {
				itemAsserted, isInnerMap := item.(map[interface{}]interface{})
				if isInnerMap {
					cleanInnerMap := cleanYml(itemAsserted)
					cleanYmlMapping[assertedKey] = cleanInnerMap
				}

			}

		}
	}
	return cleanYmlMapping
}

func CreateJsonBasedJson(filePath string) {
	fmt.Printf("Creating json file based on exising Yml config\n")
	filePathDest := filepath.Join("", filePath)
	filePathSrc := filepath.Join("", "config-sample.yml")
	file, err := os.Create(filePathDest)
	defer file.Close()
	yamlData, err := UnmarshalYAMLFile(filePathSrc)
	CheckError(err)
	//one method
	FormattedYml := cleanYml(yamlData)
	jsonOutput, err := json.MarshalIndent(FormattedYml, "", "\t")
	CheckError(err)
	file.Write(jsonOutput)
	CheckError(err)
	fmt.Println("Json data written")
}
