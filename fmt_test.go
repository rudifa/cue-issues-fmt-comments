package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/rudifa/cue-issues-fmt-comments/inproc"
)

// Test2567 is a test for issue #2567
func Test2567(t *testing.T) {
	const file = "testdata/2567-1.cue"
	err := inproc.RunCue("fmt", file)
	if err != nil {
		t.Errorf("Error running cue: %v", err)
	}
	restore(file)
}

func Test2567_2(t *testing.T) {
	const file = "testdata/2567-3.cue"
	err := inproc.RunCue("fmt", file)
	if err != nil {
		t.Errorf("Error running cue: %v", err)
	}
	restore(file)
}

func Test2274_01s(t *testing.T) {
	const file = "testdata/2274-01s.cue"
	err := inproc.RunCue("fmt", file)
	if err != nil {
		t.Errorf("Error running cue: %v", err)
	}
	restore(file)
}

func Test2274_01n(t *testing.T) {
	const file = "testdata/2274-01n.cue"
	err := inproc.RunCue("fmt", file)
	if err != nil {
		t.Errorf("Error running cue: %v", err)
	}
	restore(file)
}

// restore executes `git restore <file>`
func restore(file string) error {
	cmd := exec.Command("git", "restore", file)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing git restore: ", err)
		return err
	}
	return nil
}
