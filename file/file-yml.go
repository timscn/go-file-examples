package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	yamlContent = []byte(`---
prod: true
server_opts:
  ports: [80, 443]
  timeout: 200s
`)
)

type Configuration struct {
	IsInProductionMode bool `yaml:"prod"`
	ServerOptions      struct {
		Ports   []uint
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"server_opts"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFile() {
	fmt.Printf("Creating yml file based on pre-defined config\n")
	path := filepath.Join("", "config.yml")
	file, err := os.Create(path)
	check(err)
	defer file.Close()

	var conf Configuration

	marshalData, err := yaml.Marshal(&conf)
	check(err)
	file.Write(marshalData)
	check(err)
	fmt.Println("Yml data written")
}
