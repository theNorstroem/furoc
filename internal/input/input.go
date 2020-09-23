package input

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func GetInputYaml(specDir string, command *exec.Cmd) (error, []byte) {
	spectools := command
	spectools.Stderr = os.Stderr
	spectools.Dir = specDir

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
	return err, specYaml
}
