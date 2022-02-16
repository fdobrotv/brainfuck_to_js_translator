package external_data_test

import (
	"log"
	"os"
	"test_helper"
	"testing"
)

const (
	TestDataDir        = "../external_data/Brainfuck/testing/"
	BrainfuckExtension = ".b"
	InputExtension     = ".in"
	OutputExtension    = ".out"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestBench(t *testing.T) {
	t.Logf("%s is going to run", t.Name())

	pwd := test_helper.GetPWD(t)
	t.Logf("%s pwd", pwd)
	var testFileName = "Bench"
	brainFuckProgFile := TestDataDir + testFileName + BrainfuckExtension
	brainFuckProgOutFile := TestDataDir + testFileName + OutputExtension

	testSet := test_helper.TestSet{Program: brainFuckProgFile, Output: brainFuckProgOutFile}
	test_helper.TestTranslation(t, testSet)

	t.Logf("%s is completed", t.Name())
}

func TestBeer(t *testing.T) {
	t.Logf("%s is going to run", t.Name())

	pwd := test_helper.GetPWD(t)
	t.Logf("%s pwd", pwd)
	var testFileName = "Beer"
	brainFuckProgFile := TestDataDir + testFileName + BrainfuckExtension
	brainFuckProgOutFile := TestDataDir + testFileName + OutputExtension

	testSet := test_helper.TestSet{Program: brainFuckProgFile, Output: brainFuckProgOutFile}
	test_helper.TestTranslation(t, testSet)

	t.Logf("%s is completed", t.Name())
}
