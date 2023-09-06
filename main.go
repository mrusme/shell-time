package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/mrusme/shell-time/history"
)

type CommandStats struct {
	Command string
	Count   int64
}

type Stats struct {
	Commands map[string]CommandStats
}

func main() {
	// TODO: Filepath
	hist, err := history.New("zsh", "/home/mrus/.zsh_history")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	var stats Stats
	stats.Commands = make(map[string]CommandStats)

	numLines := hist.GetNumberOfLines()
	for i := int64(0); i < numLines; i++ {
		_, cmd, _, err := hist.GetLine(i)
		if err != nil {
			// TODO: Handle error
			continue
		}

		cmd = filepath.Base(cmd)

		cmdstat, ok := stats.Commands[cmd]
		if !ok {
			cmdstat = CommandStats{}
		}

		cmdstat.Command = cmd
		cmdstat.Count++

		stats.Commands[cmd] = cmdstat
	}

	keys := make([]string, 0, len(stats.Commands))
	for key := range stats.Commands {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return stats.Commands[keys[i]].Count < stats.Commands[keys[j]].Count
	})

	for _, cmd := range keys {
		fmt.Printf("%s: %d\n", stats.Commands[cmd].Command, stats.Commands[cmd].Count)
	}
}
