// Copyright 2023 Rudolf Farkas
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inproc

import (
	"context"

	"cuelang.org/go/cmd/cue/cmd"
)

// RunCue runs a cue command with its arguments, e.g.
// err := inproc.RunCue("eval", "testdata/sample.cue")
func RunCue(args ...string) error {
	c, _ := cmd.New(args)
	return c.Run(context.Background())
}
