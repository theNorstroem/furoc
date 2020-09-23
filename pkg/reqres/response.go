package reqres

// use this to make a reqres for furoc

import (
	"encoding/gob"
	"os"
)

type TargetFile struct {
	Filename string // full qualified name from out dir. You can start with /
	Content  []byte // the file content
}

type Response struct {
	Files []TargetFile
}

func NewResponser() *Response {
	return &Response{Files: []TargetFile{}}
}

func (r *Response) AddFile(file TargetFile) {
	r.Files = append(r.Files, file)
}

func (r *Response) SendResponse() {
	// encode and send the reqres
	encoder := gob.NewEncoder(os.Stdout)
	encoder.Encode(r)
}
