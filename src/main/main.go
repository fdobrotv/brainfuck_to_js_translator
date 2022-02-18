package main

import (
	"bytes"
	"fmt"
	"helper"
	"os"
	"translator_service"
)

const outputFile = "tmp/translated.js"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [file.bf]\n", os.Args[0])
		os.Exit(4)
	}

	progFile, err := os.Open(os.Args[1])
	defer progFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}

	inputFile, err := os.Open(os.Args[2])
	defer inputFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	prog := helper.ReadFileToByteArray(progFile)
	input := helper.ReadFileToByteArray(inputFile)

	var reader = bytes.NewReader(input)

	err = translator_service.Translate(prog, reader, outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
