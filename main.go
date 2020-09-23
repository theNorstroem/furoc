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
	commandParams := []string{"-t=test.tpl", "other"}

	files, err := subcommand.ExecuteSubcommand("./sample/furoc-gen-sample/furoc-gen-sample", specYaml, commandParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files.Files[0])

}
