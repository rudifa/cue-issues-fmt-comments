package main

import (
	"fmt"

	"github.com/rudifa/cue-issues-fmt-comments/inproc"
)

func main() {
	fmt.Println("Here we go")

	inproc.RunCue("version")

	inproc.RunCue("fmt", "testdata/sample.cue")

	inproc.RunCue("vet", "testdata/sample.cue")

	inproc.RunCue("def", "testdata/sample.cue")

	inproc.RunCue("eval", "testdata/sample.cue")

	inproc.RunCue("export", "testdata/sample.cue")

}
