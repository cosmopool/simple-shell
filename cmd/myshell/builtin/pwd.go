package builtin

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/env"
)

func executePwdCommand(args []string) {
	fmt.Println(env.SessionEnv.Pwd)
}
