package main

import (
	furoc "github.com/theNorstroem/furoc/pkg/reqres"
)

func main() {
	req := furoc.NewRequester()
	req.Fprintln(req.Parameters)

	res := furoc.NewResponser()

	// sample file
	readme := furoc.TargetFile{
		Filename: "/readme.md",
		Content:  []byte("#Test reqres"),
	}

	// build a sample reqres
	res.AddFile(readme)
	res.AddFile(readme)
	res.AddFile(readme)

	res.SendResponse()
}
