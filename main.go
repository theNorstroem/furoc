package main

import (
	"fmt"
	"github.com/theNorstroem/furoc/internal/input"
	"github.com/theNorstroem/furoc/internal/subcommand"
	"github.com/theNorstroem/furoc/pkg/parseargs"
	"github.com/theNorstroem/furoc/pkg/reqres"
	"log"
	"os"
	"os/exec"
)

func main() {
	var arglist parseargs.Arglist
	if len(os.Args) == 1 {
		// look for a .spectools config in cwd
		// if we are in a spec project and have furoc instructions there, follow them
	} else {
		// parse furoc command arguments
		arglist = parseargs.Parse()
	}

	specDir := "/Users/veith/Projects/tests/spectest"
	err, specYaml := input.GetInputYaml(specDir,
		exec.Command("/Users/veith/Projects/golang/bin/spectools", "exportAsYaml", "-f"))
	if err != nil {
		log.Fatal(err)
	}

	files := &reqres.Response{}
	for _, cmd := range arglist.Commands {
		f, err := subcommand.ExecuteSubcommand(cmd.Plugin, specYaml, cmd.Args)
		if err != nil {
			log.Fatal(err)
		}
		files = f
	}

	files = files
	fmt.Println(len(files.Files))

}
