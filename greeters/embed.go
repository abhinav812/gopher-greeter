package greeters

import "embed"

//go:embed *.txt
var embedGopherFiles embed.FS

//EmbedGopherFiles - return embedded gopher ascii-art text files.
func EmbedGopherFiles() embed.FS {
	return embedGopherFiles
}
