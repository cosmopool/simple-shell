package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	name string
	args []string
}

// Returns a [Command] from input string
func getCommandFromInput(input string) Command {
	input = strings.Trim(input, "\n")
	args := strings.Split(input, " ")
	return Command{
		name: args[0],
		args: args[1:],
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
			parseExitCommand(command.args)
		case "echo":
			parseEchoCommand(command.args)
		default:
			result := fmt.Sprintf("%s: command not found\n", command.name)
			fmt.Fprint(os.Stderr, result)
		}
	}
}
