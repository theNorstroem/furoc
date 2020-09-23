package reqres

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Request struct {
	Parameters   []string
	ParameterMap map[string]string
	AST          map[string]interface{}
}

// print what you want to the stderr console and not to the reqres
func (Request) Fprintln(i ...interface{}) {
	fmt.Fprintln(os.Stderr, i...)
}

// to enable debuging add arguments the  "debug debugfile=./sample/fullyaml.yaml"
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
