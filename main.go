package main

import (
	"fmt"
	"flag"
)

var yamlFile *string

func main() {

	yamlFile = flag.String("f", "", "YAML file")
	flag.Parse()

	// check if file exists, or exit with message (can flag handle this?)
	err := processYAML(*yamlFile)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
