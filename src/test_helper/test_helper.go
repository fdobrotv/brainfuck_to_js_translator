package test_helper

import (
	"bytes"
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
	"translator_service"
)

type TestSet struct {
	Program string
	Input   string
	Output  string
}

func TestTranslation(t *testing.T, testSet TestSet) {
	var inputBytes []byte
	if len(testSet.Input) > 0 {
		file := openFile(testSet.Input)
		inputBytes = helper.ReadFileToByteArray(file)
		defer file.Close()
	} else {
		inputBytes = nil
	}

	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)

	osOpenInput := openFile(testSet.Program)

	tempDir := t.TempDir()

	program := helper.ReadFileToByteArray(osOpenInput)
	outputFile := tempDir + "/translated.js"

	var input = bytes.NewReader(inputBytes)
	var err = translator_service.Translate(program, input, outputFile)
	if err != nil {
		os.Exit(5)
	}

	command := "node " + outputFile
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		panic(err)
	}

	output := string(data)
	t.Logf("Output is: %s", output)

	expectedOut, err := ioutil.ReadFile(testSet.Output)
	if err != nil {
		log.Fatal(err)
	}

	if string(expectedOut) != output {
		os.Exit(5)
	}
}

func GetPWD(t *testing.T) string {
	getwd, getwdErr := os.Getwd()
	if getwdErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", getwdErr)
		os.Exit(2)
	}
	t.Logf("Current test pwd is: %s", getwd)
	return getwd
}

func openFile(file string) *os.File {
	osOpenInput, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	return osOpenInput
}
