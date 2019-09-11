package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Imports struct {
	Package     string
	Subpackages []string
}
type GlideYaml struct {
	Package string
	Import  []Imports
}

func main() {
	filename := os.Args[1]
	var glide GlideYaml
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &glide)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", glide.Import[0].Package)
}
