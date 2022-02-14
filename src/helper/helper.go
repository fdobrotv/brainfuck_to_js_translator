package helper

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFileToByteArray(osOpenInput *os.File) []byte {
	bytes, err := ioutil.ReadAll(osOpenInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	return bytes
}
