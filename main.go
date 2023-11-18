package main

import (
	"fmt"

	"github.com/rudifa/cue-issues-fmt-comments/inproc"
)

func main() {
	fmt.Println("Here we go")
	inproc.RunCue("version")
	inproc.RunCue("eval", "testdata/sample.cue")
}
