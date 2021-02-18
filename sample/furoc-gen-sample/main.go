package main

import (
	"github.com/ghodss/yaml"
	furoc "github.com/theNorstroem/furoc/pkg/reqres"
	"google.golang.org/protobuf/types/pluginpb"
	"log"
)

func main() {

	// receive the request, all the stdin and mapping stuff is done here
	// if you need to debug your plugin, start it using following arguments:
	// debug debugfile=./sample/fullyaml.yaml
	// to create a debugfile
	// To create a debug file use "debugfileout=./sample/fullyaml.yaml" as an argument
	req, res := furoc.NewRequester()

	// create a responser, which can be used to add files and send the response back to furoc

	// use req.Fprintln(interface{}) if you want to print something to the console (or write to stderr)
	// you can not write to the console with fmt or log (because this goes to stdout)
	req.Fprintln("Sample plugin started")

	// the req object  contains
	//	Parameters   []string  a list of the given input parameters
	//	ParameterMap map[string]string the input parameters transformed to a map
	//	AST          AST

	for name, s := range req.AST.Services {

		// Using your own extension
		// when you have the custom extension "sampleExtension" in the service spec
		//
		//         extensions:
		//            sampleExtension:
		//                generate: sample
		// you can decode its content with furoc.DecodeExtension
		ext := &MyServiceSpecExtension{}
		furoc.DecodeExtension(s.ServiceSpec.Extensions, "sampleExtension", ext)

		// do something if generate was set in the extension
		if ext.generate {

			// build your file
			fileContent, err := yaml.Marshal(s.ServiceSpec)
			if err != nil {
				log.Fatal(err)
			}

			// if your plugin needs to call another executable, you can use commandpipe.NewCommand()

			name := "/" + name + "/" + s.ServiceSpec.Name + ".md"
			cntnt := string(fileContent)
			// create sample file
			readme := pluginpb.CodeGeneratorResponse_File{
				Name:    &name, // full qualified filename which will generated in :outputdir/
				Content: &cntnt,
			}

			// Add file to the responder
			res.AddFile(&readme)
		}
	}

	// send the response back to furoc
	res.SendResponse()
}

type MyServiceSpecExtension struct {
	generate bool `yaml:"generate"`
}
