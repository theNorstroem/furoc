package subcommand

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/theNorstroem/furoc/pkg/reqres"
	"log"
	"os"
	"os/exec"
)

func ExecuteSubcommand(command string, specYaml []byte, params []string) (files *reqres.Response, err error) {
	subProcess := exec.Command(command, params...)
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer // buffer the spectools output here
	//subProcess.Stdout = os.Stdout
	subProcess.Stdout = &b
	subProcess.Stderr = os.Stderr

	gobDecoder := gob.NewDecoder(&b)

	fmt.Println("START")                      //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
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
	fmt.Println("END") //for debug
	return &r, nil
}
