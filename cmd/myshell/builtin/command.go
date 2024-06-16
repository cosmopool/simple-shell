package builtin

import (
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/env"
)

type BuiltinCommandFunction func(args []string)

type Command struct {
	Name      string
	Args      []string
	IsBuiltin bool
	Execute   BuiltinCommandFunction
}

// Returns a [Command] from input string
func GetCommandFromInput(input string) Command {
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
		Name:      name,
		Args:      args,
		IsBuiltin: isBuiltin,
		Execute:   builtinFunction,
	}
}

// List of all builtin commands
func GetBuiltinCommands() [3]string {
	return [...]string{"exit", "echo", "type"}
}

func getCommandPath(command Command) (string, error) {
	var pathToCommand string

	for _, basePath := range env.SessionEnv.Path {
		pathToCommand = basePath + string(os.PathSeparator) + command.Name

		_, err := os.Stat(pathToCommand)
		if err != nil {
			continue
		}

		return pathToCommand, nil
	}
	return "", os.ErrNotExist
}
