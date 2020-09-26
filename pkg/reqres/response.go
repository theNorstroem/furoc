package reqres

// use this to make a reqres for furoc

import (
	"encoding/gob"
	"github.com/theNorstroem/spectools/pkg/util"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type TargetFile struct {
	Filename string // full qualified name from out dir. You can start with /
	Content  []byte // the file content
}

type Response struct {
	Files []*TargetFile
	debug bool
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
	if r.debug {
		// do the writes directly when debuging is enabled
		for _, file := range r.Files {
			if util.DirExists("debug_out") {
				fname := path.Join("debug_out", file.Filename)
				util.MkdirRelative(path.Dir(fname))
				ioutil.WriteFile(fname, file.Content, 0644)
			} else {
				log.Fatal("Dir does not exist: ", "debug_out")
			}

		}
	} else {

		// encode and send the reqres
		encoder := gob.NewEncoder(os.Stdout)
		encoder.Encode(r)
	}
}
