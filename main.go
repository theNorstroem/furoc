package main

import (
	"fmt"
	"github.com/theNorstroem/furoc/internal/input"
	"github.com/theNorstroem/furoc/internal/subcommand"
	"github.com/theNorstroem/furoc/pkg/parseargs"
	"github.com/theNorstroem/furoc/pkg/reqres"
	"github.com/theNorstroem/spectools/pkg/util"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
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

	type CMDResponse struct {
		response      *reqres.Response
		baseTargetDir string
	}
	allResponses := []CMDResponse{}
	//  for duplicate file check
	fullFilelist := map[string]bool{}

	for _, cmd := range arglist.Commands {
		r, err := subcommand.ExecuteSubcommand(cmd.Plugin, specYaml, cmd.Args)
		if err != nil {
			log.Fatal(err)
		}
		allResponses = append(allResponses, CMDResponse{
			response:      r,
			baseTargetDir: cmd.OutputDir,
		})
		// check for duplicate files
		for _, f := range r.Files {
			fname := cmd.OutputDir + "/" + f.Filename
			_, alreadyRagistred := fullFilelist[fname]
			if alreadyRagistred {
				log.Fatal(fname, " try to write same file twice")
			} else {
				fullFilelist[fname] = true
			}
		}
	}

	// Writer:

	for _, responseSet := range allResponses {
		for _, file := range responseSet.response.Files {
			if util.DirExists(responseSet.baseTargetDir) {
				fname := path.Join(responseSet.baseTargetDir, file.Filename)
				util.MkdirRelative(path.Dir(fname))
				ioutil.WriteFile(fname, file.Content, 0644)
				fmt.Println(fname)
			} else {
				log.Fatal("Dir does not exist: ", responseSet.baseTargetDir)
			}

		}
	}

}
