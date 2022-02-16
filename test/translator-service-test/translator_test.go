package translator_service_test

import (
	"log"
	"os"
	"test_helper"
	"testing"
)

const (
	TestDataDir        = "../data/"
	BrainfuckExtension = ".bf"
	InputExtension     = ".in"
	OutputExtension    = ".out"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestHelloWorld(t *testing.T) {
	t.Logf("%s is going to run", t.Name())
	var testFileName = "HelloWorld"
	brainFuckProgFile := TestDataDir + testFileName + BrainfuckExtension
	brainFuckProgInFile := ""
	brainFuckProgOutFile := TestDataDir + testFileName + OutputExtension

	testSet := test_helper.TestSet{Program: brainFuckProgFile, Input: brainFuckProgInFile, Output: brainFuckProgOutFile}
	test_helper.TestTranslation(t, testSet)

	t.Logf("%s is completed", t.Name())
}

func TestReverseInput(t *testing.T) {
	t.Logf("%s is going to run", t.Name())
	var testFileName = "ReverseInput"
	brainFuckProgFile := TestDataDir + testFileName + BrainfuckExtension
	brainFuckProgInFile := TestDataDir + testFileName + InputExtension
	brainFuckProgOutFile := TestDataDir + testFileName + OutputExtension

	testSet := test_helper.TestSet{Program: brainFuckProgFile, Input: brainFuckProgInFile, Output: brainFuckProgOutFile}
	test_helper.TestTranslation(t, testSet)

	t.Logf("%s is completed", t.Name())
}
