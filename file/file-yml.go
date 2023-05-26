package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	IsInProductionMode bool `yaml:"prod"`
	ServerOptions      struct {
		Ports   []uint
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"server_opts"`
}

func CreateYmlFile(filePath string) {
	fmt.Printf("Creating yml file based on pre-defined config\n")
	path := filepath.Join("", filePath)
	file, err := os.Create(path)
	CheckError(err)
	defer file.Close()
	var yamlConfig Configuration
	yamlFile, err := ioutil.ReadFile("config-sample.yml")
	CheckError(err)
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	CheckError(err)
	CheckError(err)
	file.Write(yamlFile)
	CheckError(err)
	fmt.Println("Yml data written")
}
