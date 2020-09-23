package main

import (
	"github.com/ghodss/yaml"
	furoc "github.com/theNorstroem/furoc/pkg/reqres"
	"log"
)

func main() {
	req := furoc.NewRequester()
	res := furoc.NewResponser()

	req.Fprintln("Start")

	// for convinience
	ast := req.AST

	for name, s := range ast.Services {
		req.Fprintln(name)
		c, err := yaml.Marshal(s.ServiceSpec)
		if err != nil {
			log.Fatal(err)
		}
		// sample file
		readme := furoc.TargetFile{
			Filename: "/" + s.ServiceSpec.Name + ".md",
			Content:  c,
		}

		// build a sample reqres
		res.AddFile(&readme)
	}

	res.SendResponse()
}
