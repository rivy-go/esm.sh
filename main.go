package main

import (
	"embed"

	"LOCAL/server"
)

//go:embed embed
//go:embed README.md
var fs embed.FS

func main() {
	server.Serve(&fs)
}
