package zsh

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var CMDRX = regexp.
	MustCompile(`(?m)(?P<env>\w+=[^[:space:]]* ){0,}(?P<cmd>[^[:space:]]*)( ){0,}(?P<args>.*)`)

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

	return history.ParseHistFile(f)
}

func (history *History) ParseHistFile(f io.Reader) error {
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

				history.lines = append(history.lines, HistoryLine{
					line:      tmpline,
					timestamp: time.Unix(t, 0),
					command:   history.getCommand(tmpsplit[1]),
					args:      history.getArgs(tmpsplit[1]),
				})
				tmpline = ""
			}
			tmpline += string(c)
		}
	}
	return nil
}

func (history *History) GetNumberOfLines() int64 {
	return int64(len(history.lines))
}

func (history *History) GetLines() []*string {
	var lines []*string
	for i := 0; i < len(history.lines); i++ {
		lines = append(lines, &history.lines[i].line)
	}
	return lines
}

func (history *History) GetLine(lineIdx int64) (time.Time, string, string, error) {
	if lineIdx >= int64(len(history.lines)) {
		return time.Now(),
			"",
			"",
			errors.New("Index out of range")
	}
	return history.lines[lineIdx].timestamp,
		history.lines[lineIdx].command,
		history.lines[lineIdx].args,
		nil
}

func (history *History) getCommand(line string) string {
	f := CMDRX.FindStringSubmatch(line)
	c := CMDRX.SubexpNames()

	for i, cn := range c {
		if cn == "cmd" {
			return filepath.Base(f[i])
		}
	}

	return "NULL"
}

func (history *History) getArgs(line string) string {
	f := CMDRX.FindStringSubmatch(line)
	c := CMDRX.SubexpNames()

	for i, cn := range c {
		if cn == "args" {
			return f[i]
		}
	}

	return ""
}
