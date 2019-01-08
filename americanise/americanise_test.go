package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const (
	inFilename       = "british.txt"
	expectedFilename = "expected_american.txt"
	actualFilename   = "actual_american.txt"
)

func TestAmericanize(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST americanize")

	path, _ := filepath.Split(os.Args[0])
	var inFile, outFile *os.File
	var err error

	fmt.Println(path)
	inFilename := filepath.Join(path, inFilename)
	if inFile, err = os.Open(inFilename); err != nil {
		t.Fatal(err)
	}
	defer inFile.Close()

	outFilename := filepath.Join(path, actualFilename)
	if outFile, err = os.Create(outFilename); err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	defer os.Remove(outFilename)

	if err := americanise(inFile, outFile); err != nil {
		t.Fatal(err)
	}

	compare(outFilename, filepath.Join(path, expectedFilename), t)
}

func compare(actual, expected string, t *testing.T) {
	if actualBytes, err := ioutil.ReadFile(actual); err != nil {
		t.Fatal(err)
	} else if expectedBytes, err := ioutil.ReadFile(expected); err != nil {
		t.Fatal(err)
	} else {
		if bytes.Compare(actualBytes, expectedBytes) != 0 {
			t.Error("actual != expected")
		}
	}
}
