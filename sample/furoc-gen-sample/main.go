package main

import (
	"encoding/gob"
	furoc "github.com/theNorstroem/furoc/pkg/response"
	"os"
)

func main() {

	res := furoc.Response{Files: []furoc.TargetFile{}}

	// sample file
	readme := furoc.TargetFile{
		Filename: "/readme.md",
		Content:  []byte("#Test"),
	}
	// build a sample response
	res.Files = append(res.Files, readme)

	// encode and send the response
	encoder := gob.NewEncoder(os.Stdout)
	encoder.Encode(res)

}
