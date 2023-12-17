package main

import (
	"fmt"
	// "go/format"
	"strings"

	// "go/parser"
	"log"
	"os/exec"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/parser"
	"github.com/davecgh/go-spew/spew"
	"github.com/rudifa/cue-issues-fmt-comments/inproc"
	"github.com/rudifa/goutil/files"
)

func main() {
	// log.Printf("spew.Config: %+v\n", spew.Config)
	spew.Config = spew.ConfigState{Indent: "    ·", DisablePointerAddresses: true}
	// spew.Config.MaxDepth = 1
	// log.Printf("spew.Config: %+v\n", spew.Config)

	// runCueFmt_2567()
	runParseAndFormat_2567()


	// runCueFmt_2672()
	// runParseAndFormat_2672()

}

func runParseAndFormat_2567() {
	const dospew = false
	// runParseAndFormat("testdata/2567.cue", dospew)
	// runParseAndFormat("testdata/2567-2.cue", dospew)
	runParseAndFormat("testdata/2567-3.cue", dospew)
}

func runParseAndFormat_2672() {
	const dospew = false
	runParseAndFormat("testdata/2672/make_tool.cue", dospew)
}

// ----------------------------------------------------------

func runParseCueString1() {

	// runParseCueString("a:1")
	// runParseCueString("abra:1 // comment")
	// runParseCueString("foo:bar:baz:123")
	runParseCueString("{foo:1, bar:2, baz:3}")

}

func runParserWithNode1() {

	sampleFile := "testdata/2567-compr+comment.cue"

	runParseAndFormat(sampleFile, true)

}

func runCueFmt_2567() {

	sampleFile := "testdata/2567-3.cue"

	runCueFmt(sampleFile)

	// sampleFile = "testdata/2567-struct+comment-pass.cue"

	// runCueFmt(sampleFile)
}


func runCueFmt_2672() {

	sampleFile := "testdata/2672/make_tool.cue"
	runCueFmt(sampleFile)
}

// ----------------------------------------------------------
func runParseAndFormat(filename string, dospew bool) {

	fmt.Println("... runParseAndFormat ----------------------------------------")
	fmt.Printf("... input: [%s]\n", filename)

	files.CatFile(filename)

	content, _ := files.ReadString(filename)

	f, err := parser.ParseFile(filename, content, parser.ParseComments)

	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	if dospew {
		fmt.Printf("... spew.Dump:\n")
		spew.Dump(f)
	}

	outbytes, err := format.Node(f)
	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	// dbs := parser.DebugStrLong("f", f)
	// fmt.Println("... intermediate parser.DebugStrLong(f):\n", dbs)

	dbs := parser.DebugStrIndent(false, "f", f)
	fmt.Println("... intermediate parser.DebugStrIndent(f):\n", dbs)

	dbs = parser.DebugStrIndent(true, "f", f)
	fmt.Println("... intermediate parser.DebugStrIndent(f):\n", dbs)

	outstring := string(outbytes)
	fmt.Println("... output format.Node(f):\n", outstring)

}

func runCueFmt(filename string) {

	log.Println("----------------------------------------")
	log.Printf("runCueFmt: [%s]\n", filename)

	log.Println("before cue fmt")
	files.CatFile(filename)

	inproc.RunCue("fmt", filename)

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

func runParseFile(filename string) {
	fmt.Println("runParseFile ----------------------------------------")

	content, _ := files.ReadString(filename)

	f, err := parser.ParseFile(filename, content, parser.ParseComments)

	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	fmt.Printf("DebugStrLong: %s\n", parser.DebugStrIndent(false, "test parser", f))
	// parser.DebugStrLong("test parser", f)

	vs := fmt.Sprintf("%#v", f)

	fmt.Println(vs)

	spew.Dump(f)

	spew.Printf("f: %+v\n", f)

	spew.Printf("f: %#v\n", f)

}

func runParseCueString(cuestring string) {
	fmt.Println("--- runParseCueString ----------------------------------------")
	fmt.Printf("... cuestring: |%s|\n", cuestring)

	f, err := parser.ParseFile("cuestring", cuestring, parser.ParseComments)

	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}
	fmt.Printf("DebugStrIndent: %s\n", parser.DebugStrIndent(false, "test parser", f))

	// vs := fmt.Sprintf("%#v", f)

	// fmt.Println(vs)

	fmt.Println("... spew.Dump:")
	spew.Dump(f)

	// fmt.Println("--- spew.Printf(f: %+v):")
	// spew.Printf("f: %+v\n", f)

	// fmt.Println("--- spew.Printf(f: %#v):")
	// spew.Printf("f: %#v\n", f)

	// func Node(node ast.Node, opt ...Option) ([]byte, error)

	outbytes, err := format.Node(f)
	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	outstring := string(outbytes)

	fmt.Println("... format.Node(f):\n", outstring)

}

func restoreFile(filename string) error {
	cmd := exec.Command("git", "checkout", "--", filename)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
