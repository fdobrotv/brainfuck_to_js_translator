package helper

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadProgram(osOpenInput *os.File) []byte {
	prog, err := ioutil.ReadAll(osOpenInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	return prog
}
