package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	yaml "github.com/goccy/go-yaml"
)

type demo struct {
	name string
	age  int
}

func main() {
	file := flag.String("f", "", "yaml file")
	flag.Parse()
	cont, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file %s %v", *file, err))
	}
	d := demo{}
	err = yaml.Unmarshal(cont, &d)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse yaml %v", err))
	}
}
