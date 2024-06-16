package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtin"
)

func writeToStderr(errorString string) {
	_, err := fmt.Fprint(os.Stderr, errorString)
	if err != nil {
		msg := fmt.Sprintf("%v ocurred when trying to write %v to STDERR.", err, errorString)
		panic(msg)
	}
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Stderr.Write([]byte(fmt.Sprint(err)))
			os.Exit(1)
		}

		// parse input command
		command := builtin.GetCommandFromInput(rawInput)
		// execute builtin command
		if command.IsBuiltin {
			command.Execute(command.Args)
			continue
		}

		// execute native commmand
		result := fmt.Sprintf("%s: command not found\n", command.Name)
		writeToStderr(result)
	}
}
