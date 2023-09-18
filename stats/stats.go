package stats

import (
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
	Commands             map[string]CommandStats
	Hours                map[int]int
	MinutesPerDay        map[string]uint16
	AverageMinutesPerDay uint16
}

type TopHourStat struct {
	Hour  int
	Count int64
}

type TopCommandStat struct {
	Command string
	Count   int64
}

func LoadStats(hist history.History) (Stats, error) {
	var stats Stats
	stats.Commands = make(map[string]CommandStats)
	stats.Hours = make(map[int]int)
	stats.MinutesPerDay = make(map[string]uint16)

	var previousTimestamp time.Time

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

		stats.Commands[cmd] = cmdstat
		stats.Hours[timestamp.Hour()]++

		if previousTimestamp.IsZero() {
			previousTimestamp = timestamp
		}
		diffDuration := timestamp.Sub(previousTimestamp)
		if diffDuration.Seconds() >= 60 {
			if diffDuration.Minutes() <= 5 {
				stats.MinutesPerDay[timestamp.Format("20060102")] += uint16(diffDuration.Minutes())
				stats.AverageMinutesPerDay += uint16(diffDuration.Minutes())
			}
			previousTimestamp = timestamp
		}
	}

	stats.AverageMinutesPerDay = stats.AverageMinutesPerDay / uint16(len(stats.MinutesPerDay))
	return stats, nil
}

func (stats *Stats) TopCommands(num int) []TopCommandStat {
	keys := make([]string, 0, len(stats.Commands))
	for key := range stats.Commands {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return stats.Commands[keys[i]].Count > stats.Commands[keys[j]].Count
	})

	var ret []TopCommandStat
	for i, cmd := range keys {
		ret = append(ret, TopCommandStat{
			Command: stats.Commands[cmd].Command,
			Count:   stats.Commands[cmd].Count,
		})
		if i == (num - 1) {
			break
		}
	}

	return ret
}

func (stats *Stats) TopHours(num int) []TopHourStat {
	hours := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 0,
	}

	sort.SliceStable(hours, func(i, j int) bool {
		return stats.Hours[i] > stats.Hours[j]
	})

	var ret []TopHourStat
	for i, hour := range hours {
		ret = append(ret, TopHourStat{
			Hour:  hour,
			Count: int64(stats.Hours[hour]),
		})
		if i == (num - 1) {
			break
		}
	}

	return ret
}
