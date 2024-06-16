package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getBuiltinCommands() [3]string {
	return [...]string{"exit", "echo", "type"}
}

type BuiltinCommandFunction func(args []string)

type Command struct {
	name      string
	args      []string
	isBuiltin bool
	execute   BuiltinCommandFunction
}

func writeToStderr(errorString string) {
	_, err := fmt.Fprint(os.Stderr, errorString)
	if err != nil {
		msg := fmt.Sprintf("%v ocurred when trying to write %v to STDERR.", err, errorString)
		panic(msg)
	}
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

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Stderr.Write([]byte(fmt.Sprint(err)))
			os.Exit(1)
		}

		command := getCommandFromInput(rawInput)
		if command.isBuiltin {
			command.execute(command.args)
		} else {
			result := fmt.Sprintf("%s: command not found\n", command.name)
			writeToStderr(result)
		}
	}
}
