/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/parser"
)

// ExportedStruct is used to unmarshal the JSON test cases
// that were extracted from cue/parser/parser_test.go
type ExportedStruct struct {
	Desc, In, Out string
}

// RunParserDataTest reads the test cases from the given file
// and runs the parser and formatter on each test case.
// This code was cloned from cue/parser/parser_test.go
// and expanded to run the formatter.
func RunParserDataTest(filename string) {
	// read from file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// read the file into a byte slice

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	log.Printf("filename: %s len: %d\n", filename, len(data))

	// unmarshal the JSON into a slice of ExportedStruct

	var exported []ExportedStruct
	err = json.Unmarshal(data, &exported)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	// convert the struct into a slice of testCases

	testCases := make([]struct{ desc, in, out string }, len(exported))
	for i, s := range exported {
		testCases[i] = struct{ desc, in, out string }{s.Desc, s.In, s.Out}
	}

	log.Printf("testCases: %d\n", len(testCases))

	// parse and format test cases

	for i, tc := range testCases {

		mode := []parser.Option{parser.AllErrors}
		if strings.Contains(tc.desc, "comments") {
			mode = append(mode, parser.ParseComments)
		}
		if strings.Contains(tc.desc, "function") {
			mode = append(mode, parser.ParseFuncs)
		}
		node, err := parser.ParseFile("input", tc.in, mode...)
		if err != nil {
			log.Printf("unexpected error: %v\n", err)
		}

		got := parser.DebugStr(node)
		if got != tc.out {
			log.Printf("i: %d\n    got  %q;\n    want %q", i, got, tc.out)
		}

		fmt.Printf("%d --------------------------------------------\n", i)
		fmt.Printf("%d --------- desc: |%s|\n", i, tc.desc)
		fmt.Printf("%d -------- input: |%s|\n", i, tc.in)
		fmt.Printf("%d - ast debugStr: |%s|\n", i, got)

		if i == 60 || i == 67 {
			log.Printf("%d ❌ skip this case to avoid panic in formatter - case needs investigation\n", i)
			continue // needs investigation
		}
		formatted, err := format.Node(node)
		if err != nil {
			log.Printf("%d unexpected error: %v\n", i, err)
		}
		fmt.Printf("%d ---- formatted: |%s|\n", i, formatted)
	}
}
