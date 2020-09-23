package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"os/exec"
)

func main() {

	spectools := exec.Command("/Users/veith/Projects/golang/bin/spectools", "exportAsYaml", "-f")
	spectools.Stderr = os.Stderr
	spectools.Dir = "/Users/veith/Projects/tests/spectest"

	var b bytes.Buffer // buffer the spectools output here
	spectools.Stdout = &b

	err := spectools.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = spectools.Wait()
	if err != nil {
		log.Fatal(err)
	}

	specYaml := b.Bytes()

	subProcess := exec.Command("simple-generator", "-t=test.tpl")
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

}
