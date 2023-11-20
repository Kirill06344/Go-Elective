package main

import (
	"embed"
)

//go:embed file.txt
var fileString string

//go:embed file.txt
var fileByte []byte

//go:embed *.hash
var folder embed.FS

func main() {
	println(fileString)
	println(string(fileByte))
	content1, _ :=
		folder.ReadFile("h1.hash")
	println(string(content1))
	content2, _ :=
		folder.ReadFile("h2.hash")
	println(string(content2))
}
