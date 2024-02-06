/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/rudifa/cuedo/runcue"
)

// Test2567 is a test for issue #2567
// it demonstrates running the regular cue fmt command in-process
func Test2567(t *testing.T) {
	const file = "testdata/2567/2567.cue"
	runcue.RunCue("fmt", file)
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
