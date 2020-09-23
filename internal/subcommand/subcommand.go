package subcommand

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/theNorstroem/furoc/pkg/response"
	"log"
	"os"
	"os/exec"
)

func ExecuteSubcommand(subProcess *exec.Cmd, specYaml []byte) (files *response.Response, err error) {
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

	var r = response.Response{}
	err = gobDecoder.Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("END") //for debug
	return &r, nil
}
