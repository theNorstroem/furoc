package commandpipe

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

type Command struct {
	process     *exec.Cmd
	Buffer      bytes.Buffer
	writer      *bufio.Writer
	commandName string // command name
	stdin       io.WriteCloser
}

func NewCommand(command string, args ...string) Command {
	c := Command{
		process: exec.Command(command, args...),
	}
	var err error
	c.stdin, err = c.process.StdinPipe()
	if err != nil {
		log.Fatal(command, err)
	}

	c.commandName = command

	c.process.Stdout = &c.Buffer
	c.process.Stderr = os.Stderr

	c.writer = bufio.NewWriter(c.stdin)

	return c
}

func (c Command) WriteToStdin(data []byte) (bytes.Buffer, error) {

	if err := c.process.Start(); err != nil { //Use start, not run
		return bytes.Buffer{}, err
	}

	c.writer.Write(data)

	c.stdin.Close()
	c.process.Wait()
	return c.Buffer, nil
}
