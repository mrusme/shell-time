package stats

import (
	"sort"
	"strconv"
	"time"

	"github.com/mrusme/shell-time/history"
)

type CommandStats struct {
	Command    string
	Count      int64
	Timestamps []time.Time
	Hours      map[string]int
}

type Stats struct {
	Commands             map[string]CommandStats
	Hours                map[string]int
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
	stats.Hours = make(map[string]int)
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
			cmdstat.Hours = make(map[string]int)
		}

		cmdstat.Command = cmd
		cmdstat.Count++
		cmdstat.Timestamps = append(cmdstat.Timestamps, timestamp)
		cmdstat.Hours[strconv.Itoa(timestamp.Hour())]++

		stats.Commands[cmd] = cmdstat
		stats.Hours[strconv.Itoa(timestamp.Hour())]++

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
	return stats.CommandsStat(num, true)
}

func (stats *Stats) LeastUsedCommands(num int) []TopCommandStat {
	return stats.CommandsStat(num, false)
}

func (stats *Stats) CommandsStat(num int, top bool) []TopCommandStat {
	keys := make([]string, 0, len(stats.Commands))
	for key := range stats.Commands {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if top {
			return stats.Commands[keys[i]].Count > stats.Commands[keys[j]].Count
		} else {
			return stats.Commands[keys[i]].Count < stats.Commands[keys[j]].Count
		}
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
	keys := make([]string, 0, len(stats.Hours))
	for key := range stats.Hours {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return stats.Hours[keys[i]] > stats.Hours[keys[j]]
	})

	var ret []TopHourStat
	for i, hour := range keys {
		h, _ := strconv.Atoi(hour)
		ret = append(ret, TopHourStat{
			Hour:  h,
			Count: int64(stats.Hours[hour]),
		})
		if i == (num - 1) {
			break
		}
	}

	return ret
}
