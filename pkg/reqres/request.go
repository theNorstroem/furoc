package reqres

import (
	"fmt"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type AST struct {
	Config            map[string]interface{} // contains the config of the spec project, this can be relevant
	InstalledServices map[string]*serviceAst.ServiceAst
	InstalledTypes    map[string]*typeAst.TypeAst
	Services          map[string]*serviceAst.ServiceAst
	Types             map[string]*typeAst.TypeAst
}

type Request struct {
	Parameters   []string
	ParameterMap map[string]string
	AST          AST
}

// print what you want to the stderr console and not to the reqres
func (Request) Fprintln(i ...interface{}) {
	fmt.Fprintln(os.Stderr, i...)
}

// Does all the input handling, and marshalling for you
//
// To enable debuging add arguments the  "debug debugfile=./sample/fullyaml.yaml"
// To create a debug file use "debugfileout=./sample/fullyaml.yaml"
func NewRequester() Request {
	req := Request{}
	req.Parameters = os.Args
	req.ParameterMap = map[string]string{}
	// make a param map for easy access
	// you still have to parse it
	for _, p := range req.Parameters[1:] {
		kv := strings.Split(p, "=")
		if len(kv) > 1 {
			req.ParameterMap[kv[0]] = kv[1]
		} else {
			req.ParameterMap[p] = p
		}
	}

	var data []byte
	var err error
	// this is for debuging your generator
	// call it with your-cmd debug debugfile=path/to/debuginput
	if _, debug := req.ParameterMap["debug"]; debug {
		var debugFile string
		if f, ok := req.ParameterMap["debugfile"]; ok {
			debugFile = f
		}
		data, err = ioutil.ReadFile(debugFile)

		if err != nil {
			log.Fatal(debugFile, err)
		}
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
		// write debugfile
		if f, ok := req.ParameterMap["debugfileout"]; ok {
			ioutil.WriteFile(f, data, 0644)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	err = yaml.Unmarshal([]byte(data), &req.AST) //reads yaml and json because json is just a subset of yaml
	if err != nil {
		log.Fatal(err)
	}

	return req
}
