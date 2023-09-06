package main

import (
	"fmt"
	"os"

	"github.com/mrusme/shell-time/history"
)

func main() {
	// TODO: Filepath
	hist, err := history.New("zsh", "/home/mrus/.zsh_history")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%v", (*hist.GetLines()[1]))
}
