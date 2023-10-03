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
	topCommands := s.TopCommands(10)
	luCommands := s.LeastUsedCommands(10)
	topHours := s.TopHours(5)

	fmt.Println()
	fmt.Println("=== YOUR TOP 10 COMMANDS ===")
	for i := 0; i < len(topCommands); i++ {
		fmt.Printf("%2d. %s (%d times)\n", (i + 1), topCommands[i].Command, topCommands[i].Count)
	}

	fmt.Println()
	fmt.Println("=== LONG FORGOTTEN COMMANDS ===")
	for i := 0; i < len(luCommands); i++ {
		fmt.Printf("%2d. %s\n", (i + 1), luCommands[i].Command)
	}

	fmt.Println()
	fmt.Println("=== MOST PRODUCTIVE HOURS ===")
	for i := 0; i < len(topHours); i++ {
		fmt.Printf("%2d. %d:00 (%d commands fired)\n", (i + 1), topHours[i].Hour, topHours[i].Count)
	}

	fmt.Println()

	// fmt.Printf("%v\n\n", s.MinutesPerDay)
	fmt.Printf("On average you ran commands on the shell for about %d minutes per day.\n\n", s.AverageMinutesPerDay)
}
