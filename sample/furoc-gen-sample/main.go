package main

import (
	"github.com/ghodss/yaml"
	furoc "github.com/theNorstroem/furoc/pkg/reqres"
	"log"
)

func main() {

	// receive the request, all the stdin and mapping stuff is done here
	// if you need to debug your plugin, use following arguments:
	// debug debugfile=./sample/fullyaml.yaml
	req := furoc.NewRequester()

	// create a responser, which can be used to add files and send the response back to furoc
	res := furoc.NewResponser()

	// use req.Fprintln(interface{}) if you want to print something to the console (or write to stderr)
	// you can not write to the console with fmt or log (because this goes to stdout)
	req.Fprintln("Sample plugin started")

	// for convinience
	ast := req.AST

	for name, s := range ast.Services {

		// build your file
		fileContent, err := yaml.Marshal(s.ServiceSpec)
		if err != nil {
			log.Fatal(err)
		}

		// if your plugin needs to call another executable, you can use commandpipe.NewCommand()

		// create sample file
		readme := furoc.TargetFile{
			Filename: "/" + name + "/" + s.ServiceSpec.Name + ".md", // full qualified filename which will generated in :outputdir/
			Content:  fileContent,                                   //[]byte with content
		}

		// Add file to the responder
		res.AddFile(&readme)
	}

	// send the response back to furoc
	res.SendResponse()
}
