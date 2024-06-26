package cmd

import (
	"github.com/aeazer/dirserver/cmd/share"
	"github.com/aeazer/dirserver/cmd/upload"
)

type Commander interface {
	Name() string
	Run() error
}

var (
	commands = []Commander{
		&share.Command{},
		&upload.Command{},
	}
)

// Register register commands
func Register() map[string]Commander {
	return register(commands)
}

func register(commands []Commander) map[string]Commander {
	commandMap := make(map[string]Commander)
	for _, cmd := range commands {
		commandMap[cmd.Name()] = cmd
	}
	return commandMap
}
