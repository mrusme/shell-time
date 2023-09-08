package main

import (
	"fmt"
	"os"

	"github.com/mrusme/shell-time/history"
	"github.com/mrusme/shell-time/stats"
)

func main() {
	// TODO: Filepath
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = os.Getenv("HOME") + "/.zsh_history"
	}

	hist, err := history.New("zsh", histfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	s, err := stats.LoadStats(hist)
	s.TopCommands(10)
	s.TopHours(5)
}
