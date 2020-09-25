package parseargs

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

func FromSpecToolsConfig() Arglist {
	configB, err := ioutil.ReadFile(".spectools")
	if err != nil {
		log.Fatal(err)
	}

	type fc struct {
		Furoc Arglist
	}
	type c struct {
		Build fc
	}

	config := c{}
	err = yaml.Unmarshal(configB, &config)
	if err != nil {
		log.Fatal(err)
	}

	a := config.Build.Furoc

	if len(a.Inputs) == 0 {
		a.Inputs = append(a.Inputs, "./")
	}

	a.Binary = os.Args[0]

	return a
}