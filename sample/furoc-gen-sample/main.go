package main

import (
	"fmt"
	furoc "github.com/theNorstroem/furoc/pkg/response"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var debug = true
	var data []byte
	var err error

	fmt.Fprintln(os.Stderr, os.Args)

	if debug {
		data, err = ioutil.ReadFile("sample/fullyaml.yaml")
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}
	y := map[string]interface{}{}
	err = yaml.Unmarshal([]byte(data), &y) //reads yaml and json because json is just a subset of yaml
	if err != nil {
		log.Fatal(err)
	}

	res := furoc.NewResponser()

	// sample file
	readme := furoc.TargetFile{
		Filename: "/readme.md",
		Content:  []byte("#Test response"),
	}

	// build a sample response
	res.AddFile(readme)
	res.AddFile(readme)
	res.AddFile(readme)

	res.SendResponse()
}
