package main

import "github.com/timscn/go-file-examples/file"

func main() {
	file.CreateYmlFile("config-yaml.out")
	file.CreateJsonBasedJson("config-json.out")
}
