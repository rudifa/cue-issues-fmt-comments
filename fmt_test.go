package main

import (
	"testing"

	"github.com/rudifa/cue-issue2567/inproc"
)

// Test2567 is a test for issue #2567
func Test2567(t *testing.T) {
	err := inproc.RunCue("fmt", "testdata/2567.cue")
	if err != nil {
		t.Errorf("Error running cue: %v", err)
	}
}

func Test2274(t *testing.T) {
	err := inproc.RunCue("fmt", "testdata/2274-1s.cue")
	if err != nil {
		t.Errorf("Error running cue: %v", err)
	}
}
