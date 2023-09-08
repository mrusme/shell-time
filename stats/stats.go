package stats

import (
	"fmt"
	"sort"
	"time"

	"github.com/mrusme/shell-time/history"
)

type CommandStats struct {
	Command    string
	Count      int64
	Timestamps []time.Time
	Hours      map[int]int
}

type Stats struct {
	Commands map[string]CommandStats
	Hours    map[int]int
}

func LoadStats(hist history.History) (Stats, error) {
	var stats Stats
	stats.Commands = make(map[string]CommandStats)
	stats.Hours = make(map[int]int)

	numLines := hist.GetNumberOfLines()
	for i := int64(0); i < numLines; i++ {
		timestamp, cmd, _, err := hist.GetLine(i)
		if err != nil {
			// TODO: Handle error
			continue
		}

		cmdstat, ok := stats.Commands[cmd]
		if !ok {
			cmdstat = CommandStats{}
			cmdstat.Hours = make(map[int]int)
		}

		cmdstat.Command = cmd
		cmdstat.Count++
		cmdstat.Timestamps = append(cmdstat.Timestamps, timestamp)
		cmdstat.Hours[timestamp.Hour()]++
		stats.Hours[timestamp.Hour()]++

		stats.Commands[cmd] = cmdstat
	}

	return stats, nil
}

func (stats *Stats) TopCommands(num int) {
	keys := make([]string, 0, len(stats.Commands))
	for key := range stats.Commands {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return stats.Commands[keys[i]].Count > stats.Commands[keys[j]].Count
	})

	for i, cmd := range keys {
		fmt.Printf("%s: %d\n", stats.Commands[cmd].Command, stats.Commands[cmd].Count)
		if i == (num - 1) {
			break
		}
	}
}

func (stats *Stats) TopHours(num int) {
	hours := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	sort.SliceStable(hours, func(i, j int) bool {
		return stats.Hours[i] > stats.Hours[j]
	})

	for i, hour := range hours {
		fmt.Printf("%d: %d\n", hour, stats.Hours[hour])
		if i == (num - 1) {
			break
		}
	}
}
