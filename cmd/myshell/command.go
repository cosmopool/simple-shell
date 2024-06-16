package main

import (
	"os"
	"strings"
)

type BuiltinCommandFunction func(args []string)

type Command struct {
	name      string
	args      []string
	isBuiltin bool
	execute   BuiltinCommandFunction
}

// Returns a [Command] from input string
func getCommandFromInput(input string) Command {
	input = strings.Trim(input, "\n")
	allArguments := strings.Split(input, " ")
	name := allArguments[0]
	args := allArguments[1:]

	isBuiltin := false
	var builtinFunction BuiltinCommandFunction
	switch name {
	case "exit":
		builtinFunction = executeExitCommand
		isBuiltin = true
	case "echo":
		builtinFunction = executeEchoCommand
		isBuiltin = true
	case "type":
		builtinFunction = executeTypeCommand
		isBuiltin = true
	}

	return Command{
		name:      name,
		args:      args,
		isBuiltin: isBuiltin,
		execute:   builtinFunction,
	}
}

func getBuiltinCommands() [3]string {
	return [...]string{"exit", "echo", "type"}
}

func getCommandPath(command Command) (string, error) {
	var pathToCommand string

	for _, basePath := range environment.path {
		pathToCommand = basePath + string(os.PathSeparator) + command.name

		_, err := os.Stat(pathToCommand)
		if err != nil {
			continue
		}

		return pathToCommand, nil
	}
	return "", os.ErrNotExist
}
