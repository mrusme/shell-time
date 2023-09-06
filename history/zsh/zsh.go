package zsh

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type HistoryLine struct {
	line      string
	timestamp time.Time
	command   string
	args      string
}

type History struct {
	lines []HistoryLine
}

func (history *History) Initialize(historyFile string) error {
	f, err := os.Open(historyFile)
	if err != nil {
		return err
	}

	var tmpline string
	r := bufio.NewReader(f)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		} else {
			if c == ':' &&
				len(tmpline) > 0 &&
				tmpline[len(tmpline)-1:] == string('\n') {
				t, terr := strconv.ParseInt(tmpline[2:12], 10, 64)
				if terr != nil {
					return terr
				}
				tmpsplit := strings.SplitN(tmpline, ";", 2)
				cmdsplit := strings.SplitN(tmpsplit[1], " ", 2)
				cmdargs := ""
				if len(cmdsplit) > 1 {
					cmdargs = cmdsplit[1]
				}

				history.lines = append(history.lines, HistoryLine{
					line:      tmpline,
					timestamp: time.Unix(t, 0),
					command:   cmdsplit[0],
					args:      cmdargs,
				})
				tmpline = ""
			}
			tmpline += string(c)
		}
	}
	return nil
}

func (history *History) GetLines() []*string {
	var lines []*string
	for i := 0; i < len(history.lines); i++ {
		lines = append(lines, &history.lines[i].line)
	}
	return lines
}
