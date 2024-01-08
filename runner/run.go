/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package runner

import (
	"fmt"
	"os"

	// "go/format"
	"strings"

	// "go/parser"
	"log"
	"os/exec"

	"cuelang.org/go/cmd/cuedo/util"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/parser"
	"github.com/davecgh/go-spew/spew"
	"github.com/rudifa/cuedo-fmt/inproc"
	"github.com/rudifa/goutil/files"
)

func RunParseAndFormat(file string) {

	runParseAndFormat(file)
}

// ----------------------------------------------------------
func runParseAndFormat(filename string) {

	fmt.Println("••• runParseAndFormat ----------------------------------------")

	fmt.Printf("••• input: [%s]\n", filename)
	files.CatFile(filename)

	content, _ := files.ReadString(filename)

	options := []parser.Option{parser.ParseComments}
	if os.Getenv("CUEDO_PARSER_TRACE") != "" {
		options = append(options, parser.Trace)
	}

	fmt.Printf("••• parse file: [%s]\n", filename)
	f, err := parser.ParseFile(filename, content, options...)

	if err != nil {
		log.Printf("unexpected error: %v\n", err)
	}

	// dospew := true
	// if dospew {
	// 	fmt.Printf("••• spew.Dump:\n")
	// 	spew.Dump(f)
	// 	spd := spew.Sdump(f)
	// 	fmt.Printf("••• spew.Sdump:\n%s\n", spd)
	// }

	if os.Getenv("CUEDO_PARSER_DEBUG_STR") != "" {
		dbs := util.DebugStr(f)
		fmt.Println("••• parser out DebugStr(f):\n", dbs)
	}

	if os.Getenv("CUEDO_AST_TREE") != "" {
		debugAstStr := util.DebugAstStr(f)

		fmt.Println("••• parser out DebugAstStr(f):\n", debugAstStr)
	}

	fmt.Printf("••• format.Node(f)\n")
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
	fmt.Println("••• output format.Node(f):\n", outstring)

}

// ----------------------------------------------------------
// obsolete

func runParseAndFormat_2567(file string) {
	runParseAndFormat(file)
}

func runParseAndFormat_2274() {
	// runParseAndFormat("testdata/2274/2274-3.cue", dospew)
	runParseAndFormat("testdata/2274/2274-4.cue")
}

func runParseAndFormat_2672() {
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
