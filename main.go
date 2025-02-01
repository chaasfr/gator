package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	s, err := InitState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmds := InitCmds()

	cmds.register("login", HandlerLogin)
	cmds.register("register", HandlerRegister)
	cmds.register("reset", HandlerReset)

	osArgs := os.Args
	if len(osArgs) < 2 {
		fmt.Printf("error: not enough argument, received %s\n", osArgs)
		os.Exit(1)
	}

	command := Command{
		osArgs[1],
		osArgs[2:],
	}
	if err := cmds.run(s, command); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
