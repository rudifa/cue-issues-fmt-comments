/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/rudifa/cuedo-fmt/inproc"
)

// Test2567 is a test for issue #2567
func Test2567(t *testing.T) {
	const file = "testdata/2567-1.cue"
	inproc.RunCue("fmt", file)
	restore(file)
}

func Test2567_2(t *testing.T) {
	const file = "testdata/2567-3a.cue"
	inproc.RunCue("fmt", file)
	restore(file)
}

func Test2274_01s(t *testing.T) {
	const file = "testdata/2274-01s.cue"
	inproc.RunCue("fmt", file)
	restore(file)
}

func Test2274_01n(t *testing.T) {
	const file = "testdata/2274-01n.cue"
	inproc.RunCue("fmt", file)
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
