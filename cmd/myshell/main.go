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

type Command struct {
	name string
	args []string
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

	return Command{
		name: name,
		args: args,
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
		switch command.name {
		case "exit":
			executeExitCommand(command.args)
		case "echo":
			executeEchoCommand(command.args)
		case "type":
			executeTypeCommand(command.args)
		default:
			result := fmt.Sprintf("%s: command not found\n", command.name)
			writeToStderr(result)
		}
	}
}
