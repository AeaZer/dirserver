package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/aeazer/dirserver/cmd"
	"github.com/aeazer/dirserver/utils/color"
)

const (
	modeShare  = "share"
	modeUpload = "upload"
	helpMark   = "-h"
)
func main() {
	if len(os.Args) == 1 || !slices.Contains([]string{modeShare, modeUpload, helpMark}, os.Args[1]) {
		fmt.Printf("unknown command mod, you can use dirserver -h to show usage.\n")
		os.Exit(1)
	}
    // Register commands
	cmdMap := cmd.Register()
	if os.Args[1] == helpMark {
		fmt.Println(color.BlueDA.Dyeing("GitHub: https://github.com/aeazer/dirserver\n"))
		fmt.Printf("%s:\n  %s [command] [-subcommand]\n", color.BlueDA.Dyeing("Usage"), color.GreenDA.Dyeing("dirserver"))
        // Display commands' help information
		for _, commander := range cmdMap {
			_ = commander.Run()
		}
		os.Exit(0)
	}
    // Run specified command
	err := cmdMap[os.Args[1]].Run()
	if err != nil {
		log.Fatalf("command run occurred error: %v\n", err)
	}
}
