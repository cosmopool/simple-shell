package builtin

import (
	"fmt"
	"os"
	"strconv"
)

func executeExitCommand(args []string) {
	var exitCode int

	if len(args) == 0 {
		os.Exit(0)
	}

	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(exitCode)
	}
	os.Exit(exitCode)
}
