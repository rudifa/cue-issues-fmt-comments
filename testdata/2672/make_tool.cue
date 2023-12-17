package make_tool

import (
	"tool/exec"
)

// displays help for available commands
command: help: {
	show: exec.Run & {cmd: [ "sh", "-c", "cue help cmd | grep -A2 'Available Commands:'"]}
}
