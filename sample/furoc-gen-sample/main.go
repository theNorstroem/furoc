package main

import (
	furoc "github.com/theNorstroem/furoc/pkg/response"
)

func main() {

	res := furoc.NewResponser()

	// sample file
	readme := furoc.TargetFile{
		Filename: "/readme.md",
		Content:  []byte("#Test response"),
	}

	// build a sample response
	res.AddFile(readme)
	res.AddFile(readme)
	res.AddFile(readme)

	res.SendResponse()
}
