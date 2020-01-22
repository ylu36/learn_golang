package main

import (
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	openFile(filename)
}

func openFile(filename string) {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
