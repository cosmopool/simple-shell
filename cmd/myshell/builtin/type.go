package builtin

import (
	"fmt"
	"os"
)

func executeTypeCommand(args []string) {
	if len(args) < 1 {
		return
	}

	// check if is a builtin command
	command := args[0]
	isBuiltin := false
	for _, builtinCommand := range getBuiltinCommands() {
		if command == builtinCommand {
			isBuiltin = true
			break
		}
	}

	if isBuiltin {
		fmt.Fprintln(os.Stdout, command+" is a shell builtin")
	} else {
		fmt.Fprintln(os.Stderr, command+": not found")
	}
}
