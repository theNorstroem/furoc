package subcommand

import (
	"bufio"
	"fmt"
	"github.com/theNorstroem/furoc/pkg/response"
	"log"
	"os"
	"os/exec"
)

func ExecuteSubcommand(subProcess *exec.Cmd, specYaml []byte) (files response.Response, err error) {
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create("./out.txt")
	//subProcess.Stdout = os.Stdout
	subProcess.Stdout = outfile
	subProcess.Stderr = os.Stderr

	fmt.Println("START")                      //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}

	writer := bufio.NewWriter(stdin)
	writer.Write(specYaml)

	stdin.Close()
	subProcess.Wait()
	fmt.Println("END") //for debug
	return response.Response{}, nil
}
