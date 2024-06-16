package main

import (
	"fmt"
	"os"
)

func executeTypeCommand(args []string) {
	if len(args) < 1 {
		return
	}

	// check if is a builtin commandToCheck
	commandToCheck := args[0]
	isBuiltin := false
	for _, builtinCommand := range getBuiltinCommands() {
		if commandToCheck == builtinCommand {
			isBuiltin = true
			break
		}
	}

	if isBuiltin {
		fmt.Fprintln(os.Stdout, commandToCheck+" is a shell builtin")
		return
	}

	// check non-builtin command
	commandPath, err := getCommandPath(Command{name: commandToCheck})
	if err != nil {
		fmt.Fprintln(os.Stderr, commandPath+": not found")
	}
}
