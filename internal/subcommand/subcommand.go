package subcommand

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"github.com/theNorstroem/furoc/pkg/reqres"
	"log"
	"os"
	"os/exec"
)

func ExecuteSubcommand(command string, specYaml []byte, params []string) (files *reqres.Response, err error) {
	if command == "" {
		log.Fatal("plugin not defined")
	}
	subProcess := exec.Command(command, params...)
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		log.Fatal(command, err)
	}

	var b bytes.Buffer
	subProcess.Stdout = &b
	subProcess.Stderr = os.Stderr

	gobDecoder := gob.NewDecoder(&b)

	if err = subProcess.Start(); err != nil { //Use start, not run
		log.Fatal(command, err)
	}

	writer := bufio.NewWriter(stdin)
	writer.Write(specYaml)

	stdin.Close()
	subProcess.Wait()

	var r = reqres.Response{}
	err = gobDecoder.Decode(&r)
	if err != nil {
		log.Fatal(err)
	}

	return &r, nil
}
