package translator_service

import (
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

const (
	TEST_DATA_DIR       = "../data/"
	BRAINFUCK_EXTENSION = ".bf"
	INPUT_EXTENSION     = ".in"
	OUTPUT_EXTENSION    = ".out"
)

type TestSet struct {
	program string
	input   string
	output  string
}

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestHelloWorld(t *testing.T) {
	t.Logf("%s is going to run", t.Name())
	var testFileName = "HelloWorld"
	brainFuckProgFile := TEST_DATA_DIR + testFileName + BRAINFUCK_EXTENSION
	brainFuckProgInFile := ""
	brainFuckProgOutFile := TEST_DATA_DIR + testFileName + OUTPUT_EXTENSION

	testSet := TestSet{brainFuckProgFile, brainFuckProgInFile, brainFuckProgOutFile}
	testTranslation(t, testSet)

	t.Logf("%s is completed", t.Name())
}

func TestReverseInput(t *testing.T) {
	t.Logf("%s is going to run", t.Name())
	var testFileName = "ReverseInput"
	brainFuckProgFile := TEST_DATA_DIR + testFileName + BRAINFUCK_EXTENSION
	brainFuckProgInFile := TEST_DATA_DIR + testFileName + INPUT_EXTENSION
	brainFuckProgOutFile := TEST_DATA_DIR + testFileName + OUTPUT_EXTENSION

	testSet := TestSet{brainFuckProgFile, brainFuckProgInFile, brainFuckProgOutFile}
	testTranslation(t, testSet)

	t.Logf("%s is completed", t.Name())
}

func testTranslation(t *testing.T, testSet TestSet) {
	if len(testSet.input) > 0 {

	}

	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)

	osOpenInput, err := openFile(testSet.program)

	tempDir := t.TempDir()

	program := helper.ReadProgram(osOpenInput)
	outputFile := tempDir + "/translated.js"
	err = translator_service.Translate(program, outputFile)
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

	expectedOut, err := ioutil.ReadFile(testSet.output)
	if err != nil {
		log.Fatal(err)
	}

	if string(expectedOut) != output {
		os.Exit(5)
	}
}

func getPWD(t *testing.T) string {
	getwd, getwdErr := os.Getwd()
	if getwdErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", getwdErr)
		os.Exit(2)
	}
	t.Logf("Current test pwd is: %s", getwd)
	return getwd
}

func openFile(file string) (*os.File, error) {
	osOpenInput, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	return osOpenInput, err
}
