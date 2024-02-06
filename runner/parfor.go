// Package runner implements functions that call CUE parser dd formatter functions to help in debugging them.
package runner

/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/parser"
	"github.com/davecgh/go-spew/spew"
	"github.com/rudifa/cuedo/runcue"
	"github.com/rudifa/goutil/ffmt"
	"github.com/rudifa/goutil/files"
)

// RunParseAndFormat runs the parser and formatter on the given file
func RunParseAndFormat(file string) {
	runParseAndFormat(file)
}

// ----------------------------------------------------------
func runParseAndFormat(filename string) {

	const threeBlack = "•••"
	const threeWhite = "○○○"
	fmt.Println("••• runParseAndFormat ----------------------------------------")

	fmt.Println("•••", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Printf("••• input: %s\n", filename)
	files.CatFile(filename)

	content, _ := files.ReadString(filename)


	parseComments := []bool{true}
	if os.Getenv("CUEDO_PROCESS_BOTH_WITH_AND_WITHOUT_COMMENTS") != "" {
		parseComments = append(parseComments, false)
	}

	for _, withComments := range parseComments {
		options := []parser.Option{}

		var three string
		if withComments {
			options = append(options, parser.ParseComments)
			three = threeBlack
		} else {
			three = threeWhite
		}

		if os.Getenv("CUEDO_PARSER_TRACE") != "" {
			options = append(options, parser.Trace)
		}

		fmt.Printf(three + " parse file: %s\n", filename)
		f, err := parser.ParseFile(filename, content, options...)
		if err != nil {
			log.Printf("unexpected error: %v\n", err)
		}

		if os.Getenv("CUEDO_PARSER_DEBUG_STR") != "" {
			dbs := parser.DebugStr(f)
			dbs, _ = ffmt.IndentNestedBrackets(dbs, "<>", "  ")
			fmt.Println(three + " parser out DebugStr(f):\n", dbs)
		}

		if os.Getenv("CUEDO_AST_TREE") != "" {
			debugAstStr := parser.DebugAstTree(f)

			fmt.Println(three + " parser out DebugAstTree(f):\n", debugAstStr)
		}

		if os.Getenv("CUEDO_AST_NODE_SPEW") != "" {
			fmt.Println(three + " parser out spew.Dump(f):")
			spew.Dump(f)
		}

		fmt.Printf(three + " format.Node(f)\n")
		outbytes, err := format.Node(f)
		if err != nil {
			log.Printf("unexpected error: %v\n", err)
		}

		// dbs := parser.DebugStrLong("f", f)
		// fmt.Println("••• parser.DebugStrLong(f):\n", dbs)

		// dbs := parser.DebugStrIndent(false, "f", f)
		// fmt.Println("••• parser.DebugStrIndent(f):\n", dbs)

		// dbs = parser.DebugStrIndent(true, "f", f)
		// fmt.Println("••• parser.DebugStrIndent(f):\n", dbs)

		outstring := string(outbytes)
		fmt.Println(three + " output format.Node(f):\n", outstring)
	}
}

// ----------------------------------------------------------
// obsolete

func runParseAndFormat2567(file string) {
	runParseAndFormat(file)
}

func runParseAndFormat2274() {
	// runParseAndFormat("testdata/2274/2274-3.cue", dospew)
	runParseAndFormat("testdata/2274/2274-4.cue")
}

func runParseAndFormat2672() {
	runParseAndFormat("testdata/2672/make_tool.cue")
}

func runParseCueString1() {
	// runParseCueString("a:1")
	// runParseCueString("abra:1 // comment")
	// runParseCueString("foo:bar:baz:123")
	runParseCueString("{foo:1, bar:2, baz:3}")
}

func runParserWithNode1() {
	sampleFile := "testdata/2567-compr+comment.cue"

	runParseAndFormat(sampleFile)
}

func runCueFmt2567() {
	sampleFile := "testdata/2567-3.cue"

	runCueFmt(sampleFile)

	// sampleFile = "testdata/2567-struct+comment-pass.cue"

	// runCueFmt(sampleFile)
}

func runCueFmt2672() {
	sampleFile := "testdata/2672/make_tool.cue"
	runCueFmt(sampleFile)
}

func runCueFmt(filename string) {
	log.Println("----------------------------------------")
	log.Printf("runCueFmt: [%s]\n", filename)

	log.Println("before cue fmt")
	files.CatFile(filename)

	runcue.RunCue("fmt", filename)

	log.Println("after cue fmt")
	files.CatFile(filename) // may be  modified

	// restore file `filename` from the local git repo
	err := restoreFile(filename)
	if err != nil {
		log.Printf("restoreFile: unexpected error: %v\n", err)
	}
}

func runCueCommands(filename string) {
	log.Println("----------------------------------------")
	log.Printf("cueTestCase: [%s]\n", filename)

	files.CatFile(filename)

	// runcue.RunCue("version")

	{
		copyFilename := strings.Replace(filename, ".cue", ".copy.cue", -1)
		_ = files.CopyFile(filename, copyFilename)

		runcue.RunCue("fmt", copyFilename)

		files.CatFile(copyFilename) // may be  modified
		files.RemoveFileIfExists(copyFilename)
	}

	// runcue.RunCue("vet", filename)

	// runcue.RunCue("def", filename)

	// runcue.RunCue("eval", filename)

	// runcue.RunCue("export", filename)
}

func runParseFile(filename string) {
	fmt.Println("runParseFile ----------------------------------------")

	content, _ := files.ReadString(filename)

	f, err := parser.ParseFile(filename, content, parser.ParseComments)
	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	// fmt.Printf("DebugStrLong: %s\n", parser.DebugStrIndent(false, "test parser", f))
	// parser.DebugStrLong("test parser", f)

	vs := fmt.Sprintf("%#v", f)

	fmt.Println(vs)

	spew.Dump(f)

	spew.Printf("f: %+v\n", f)

	spew.Printf("f: %#v\n", f)
}

func runParseCueString(cuestring string) {
	fmt.Println("--- runParseCueString ----------------------------------------")
	fmt.Printf("••• cuestring: |%s|\n", cuestring)

	f, err := parser.ParseFile("cuestring", cuestring, parser.ParseComments)
	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}
	// fmt.Printf("DebugStrIndent: %s\n", parser.DebugStrIndent(false, "test parser", f))

	// vs := fmt.Sprintf("%#v", f)

	// fmt.Println(vs)

	fmt.Println("••• spew.Dump:")
	spew.Dump(f)

	// fmt.Println("--- spew.Printf(f: %+v):")
	// spew.Printf("f: %+v\n", f)

	// fmt.Println("--- spew.Printf(f: %#v):")
	// spew.Printf("f: %#v\n", f)

	// func Node(node ast.Node, opt •••Option) ([]byte, error)

	outbytes, err := format.Node(f)
	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	outstring := string(outbytes)

	fmt.Println("••• format.Node(f):\n", outstring)
}

func restoreFile(filename string) error {
	cmd := exec.Command("git", "checkout", "--", filename)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
