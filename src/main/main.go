package main

import (
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

	osOpenInput, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}

	prog := helper.ReadProgram(osOpenInput)
	err = translator_service.Translate(prog, outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
