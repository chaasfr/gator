package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmdsMap map[string]func(*State, Command) error
}


func (c *Commands) register(name string, f func(*State, Command) error) {
	c.cmdsMap[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.cmdsMap[cmd.name]
	if !ok {
		return fmt.Errorf("command not implemented: %s", cmd.name)
	}
	return f(s, cmd)
}

func InitCmds() *Commands {
	var cmds Commands
	cmds.cmdsMap = make(map[string]func(*State, Command) error)
	return &cmds
}
