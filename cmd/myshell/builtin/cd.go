package builtin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/env"
)

func executeCdCommand(args []string) {
	if len(args) < 1 {
		return
	}

	targetDir := args[0]
	isHomeDir := rune(targetDir[0]) == '~'
	if isHomeDir {
		targetDir = env.SessionEnv.Home
	}

	isRelativePath := rune(targetDir[0]) == '.'
	if isRelativePath {
		targetDir = filepath.Join(env.SessionEnv.Pwd, targetDir)
	}

	_, err := os.Stat(targetDir)
	if err != nil {
		msg := fmt.Sprintf("cd: %s: No such file or directory\n", targetDir)
		fmt.Fprint(os.Stderr, msg)
	}

	err = env.SessionEnv.SetEnv(env.PWD, targetDir)
	if err != nil {
		panic(err)
	}
}
