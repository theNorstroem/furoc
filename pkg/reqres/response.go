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
	Files []*TargetFile
}

// Creates a responser which holds the files you want to send back to protoc.
func NewResponser() *Response {
	return &Response{Files: []*TargetFile{}}
}

// Add a file to the responser, duplicate filename checks are done in furoc.
func (r *Response) AddFile(file *TargetFile) {
	r.Files = append(r.Files, file)
}

// Send the encoded message response back to furoc
func (r *Response) SendResponse() {
	// encode and send the reqres
	encoder := gob.NewEncoder(os.Stdout)
	encoder.Encode(r)
}
