package builtin

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/env"
)

func executeCdCommand(args []string) {
	if len(args) < 1 {
		return
	}

	nextDir := args[0]
	_, err := os.Stat(nextDir)
	if err != nil {
		msg := fmt.Sprintf("cd: %s: No such file or directory\n", nextDir)
		fmt.Fprint(os.Stderr, msg)
	}

	err = env.SessionEnv.SetEnv(env.PWD, nextDir)
	if err != nil {
		panic(err)
	}
}
