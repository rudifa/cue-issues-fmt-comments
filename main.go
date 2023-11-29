package main

import (
	"fmt"
	"strings"

	// "go/parser"
	"log"

	"cuelang.org/go/cue/parser"

	"github.com/davecgh/go-spew/spew"
	"github.com/rudifa/cue-issues-fmt-comments/inproc"
	"github.com/rudifa/goutil/files"
)

func main() {
	// fmt.Println("Here we go")

	// testParseFile("testdata/sample.cue")

	// cueTestCase("testdata/sample.cue")

	// sampleFile := "testdata/sample.cue"
	// sampleFile := "testdata/2567-3a.cue"
	sampleFile := "testdata/fbb-c.cue"

	cueTestCase(sampleFile)

	testParser(sampleFile)

}

func cueTestCase(filename string) {

	log.Println("----------------------------------------")
	log.Printf("cueTestCase: [%s]\n", filename)

	files.CatFile(filename)

	// inproc.RunCue("version")

	{
		copyFilename := strings.Replace(filename, ".cue", ".copy.cue", -1)
		_ = files.CopyFile(filename, copyFilename)

		inproc.RunCue("fmt", copyFilename)

		files.CatFile(copyFilename) // may be  modified
		files.RemoveFileIfExists(copyFilename)
	}

	// inproc.RunCue("vet", filename)

	// inproc.RunCue("def", filename)

	// inproc.RunCue("eval", filename)

	// inproc.RunCue("export", filename)
}

func testParser(filename string) {

	log.Println("test parser.ParseFile ----------------------------------------")
	log.Printf("testParser: [%s]\n", filename)

	files.CatFile(filename)

	content := files.ReadFile(filename)


	f, err := parser.ParseFile(filename, content, parser.ParseComments)

	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}
	// parser.LogDebugStr("spewParser", f)
	// fmt.Println()

	// fmt.Println("1 spew.Printf '%+v'")
	// spew.Printf("f: %+v\n", f)

	// fmt.Println("2 spew.Printf '%#v'")
	// spew.Printf("f: %#v\n", f)

	fmt.Println("3 spew.Dump")
	spew.Dump(f)

}

func testParseFile(filename string) {
	fmt.Println("testParseFile ----------------------------------------")

	content := files.ReadFile(filename)

	f, err := parser.ParseFile(filename, content, parser.ParseComments)

	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}
	parser.LogDebugStr("test parser", f)

	vs := fmt.Sprintf("%#v", f)

	fmt.Println(vs)

	spew.Dump(f)

	spew.Printf("f: %+v\n", f)

	spew.Printf("f: %#v\n", f)

}
