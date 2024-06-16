package builtin

import "strings"

func getBuiltinCommands() [3]string {
	return [...]string{"exit", "echo", "type"}
}

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
