package main

import (
	"fmt"
	"github.com/theNorstroem/furoc/internal/input"
	"github.com/theNorstroem/furoc/internal/subcommand"
	"log"
	"os/exec"
)

func main() {
	specDir := "/Users/veith/Projects/tests/spectest"
	err, specYaml := input.GetInputYaml(specDir,
		exec.Command("/Users/veith/Projects/golang/bin/spectools", "exportAsYaml", "-f"))
	if err != nil {
		log.Fatal(err)
	}

	//subProcess := exec.Command("simple-generator", "-t=test.tpl")
	subProcess := exec.Command("./furoc-gen-sample")
	files, err := subcommand.ExecuteSubcommand(subProcess, specYaml)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)

}
