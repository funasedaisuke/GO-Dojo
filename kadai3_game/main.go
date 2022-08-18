package main

import (
	"fmt"
	"main/game"
	"os"
)

func main() {
	if err := recover(); err != nil {
		fmt.Fprintf(os.Stdout, "Error:\n%s\n", err)
		os.Exit(1)
	}
	cli := game.Game{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
