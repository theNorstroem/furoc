package response

type TargetFile struct {
	Filename string // full qualified name from out dir. You can start with /
	Content  []byte // the file content
}

type Response struct {
	Files []TargetFile
}
