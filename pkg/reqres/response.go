package reqres

// use this to make a reqres for furoc

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
)

type Response pluginpb.CodeGeneratorResponse

// Creates a responser which holds the files you want to send back to protoc.
func NewResponser() *Response {
	return &Response{
		File: []*pluginpb.CodeGeneratorResponse_File{},
	}
}

// Add a file to the responser, duplicate filename checks are done in furoc.
func (r *Response) AddFile(file *pluginpb.CodeGeneratorResponse_File) {
	r.File = append(r.File, file)
}

// Send the encoded message response back to furoc
func (r *Response) SendResponse() {
	// encode and send the reqres
	res := pluginpb.CodeGeneratorResponse(*r)
	marshalled, err := proto.Marshal(&res)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
