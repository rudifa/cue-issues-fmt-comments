/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

// Package cmd implements the cuedo commands
package cmd

import (
	"fmt"

	//	"github.com/rudifa/cuedo/runner"

	"github.com/rudifa/cuedo/runner"
	"github.com/spf13/cobra"
)

var filename = "testdata/parser_test.json"

// formatFromJSONCmd represents the formatFromJson command
var formatFromJSONCmd = &cobra.Command{
	Use:     "formatFromJson",
	Short:   "Reads cue test cases from json and parses and formats them",
	Long:    `Reads cue test cases from json and parses and formats them`,
	Aliases: []string{"ffj"},
	Run: func(cmd *cobra.Command, args []string) {
		// Get the filename from the command arguments
		// if len(args) != 1 {
		// 	log.Fatalf("Please provide a filename as argument")
		// }
		// filename := args[0]

		// filename := "testdata/parser_test.json"

		fmt.Println("formatFromJson filename", filename)

		runner.RunParserDataTest(filename)
	},
}

func init() {
	rootCmd.AddCommand(formatFromJSONCmd)
}

// // // to be moved into a separate file; untested
// type ExportedStruct struct {
// 	Desc, In, Out string
// }

// func testParserData(filename string) {
// 	// read from file
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %s", err)
// 	}
// 	defer file.Close()

// 	// read the file into a byte slice
// 	data, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatalf("Failed to read file: %s", err)
// 	}

// 	log.Printf("filename: %s len: %d\n", filename, len(data))

// 	// unmarshal the JSON encoding of the exported struct
// 	var exported []ExportedStruct
// 	err = json.Unmarshal(data, &exported)
// 	if err != nil {
// 		log.Fatalf("Failed to unmarshal JSON: %s", err)
// 	}

// 	// convert the struct into a slice of testCases
// 	testCases := make([]struct{ desc, in, out string }, len(exported))
// 	for i, s := range exported {
// 		testCases[i] = struct{ desc, in, out string }{s.Desc, s.In, s.Out}
// 	}

// 	log.Printf("testCases: %d\n", len(testCases))

// 	// testParser(testCases)
// 	for i, tc := range testCases {
// 		mode := []parser.Option{parser.AllErrors}
// 		if strings.Contains(tc.desc, "comments") {
// 			mode = append(mode, parser.ParseComments)
// 		}
// 		if strings.Contains(tc.desc, "function") {
// 			mode = append(mode, parser.ParseFuncs)
// 		}
// 		node, err := parser.ParseFile("input", tc.in, mode...)
// 		if err != nil {
// 			log.Printf("unexpected error: %v\n", err)
// 		}

// 		got := util.DebugStr(node)
// 		if got != tc.out {
// 			log.Printf("i: %d\n    got  %q;\n    want %q", i, got, tc.out)
// 		}
// 		formatted := format.Node(node)
// 		// if formatted != tc.out {
// 		// 	t.Errorf("\ngot  %q;\nwant %q", formatted, tc.out)
// 		// }
// 		fmt.Printf("i: %d formatted: %s\n", i, formatted)

// 	}
// }
