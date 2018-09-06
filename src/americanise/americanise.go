package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var britishAmerican = "british-american.txt"

func init() {
	dir := filepath.Dir(os.Args[0])
	britishAmerican = filepath.Join(dir, britishAmerican)
}

func setupLog() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "01-01-1970 23:59:59 GMT"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

func filenamesFromCommandLine() (inFilename string, outFilename string, err error) {
	if len(os.Args) != 3 {
		return "", "", fmt.Errorf("Invalid / Missing arguments")
	}

	// infileName and outfileName cannot be same
	if os.Args[1] == os.Args[2] {
		return "", "", fmt.Errorf("\ninputFilePath and outputFilePath cannot be same (%s)", os.Args[1])
	}

	return os.Args[1], os.Args[2], nil
}

func makeReplacerFunction(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	text := string(rawBytes)

	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}

func americanise(inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string
	if replacer, err = makeReplacerFunction(britishAmerican); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil // not really an error
			eof = true
		} else if err != nil {
			return err
		}
		line = wordRx.ReplaceAllStringFunc(line, replacer)
		log.Infof("Converted line: %s", line)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	setupLog()

	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("\nusage: %s <inputFilePath> <outputFilePath>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	inFile, outFile := os.Stdin, os.Stdout

	log.Infof("%s START", filepath.Base(os.Args[0]))
	log.Infof("INPUT FILE = [%s], OUTPUT FILE = [%s]", inFilename, outFilename)

	if len(strings.TrimSpace(inFilename)) > 0 {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}
	log.Info("Successfully opened input file: " + inFile.Name())

	if len(strings.TrimSpace(outFilename)) > 0 {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}
	log.Info("Successfully opened output file: " + outFile.Name())

	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}
