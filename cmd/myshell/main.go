package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtin"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/env"
)

func main() {
	env.SetEnvironment()

	for {
		writeToStdout("$ ")
		// wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			writeToStderr(err.Error())
			continue
		}

		command := builtin.GetCommandFromInput(input)
		// execute builtin command
		if command.IsBuiltin {
			command.Execute(command.Args)
			continue
		}

		// execute non-builtin command
		commandPath, err := builtin.GetCommandPath(command)
		if err != nil {
			result := fmt.Sprintf("%s: command not found\n", command.Name)
			writeToStderr(result)
			continue
		}

		cmd := exec.Command(commandPath, command.Args...)
		cmd.Stdout = os.Stdout

		err = cmd.Run()
		if err != nil {
			writeToStderr(err.Error())
		}
	}
}
