package zsh

import (
	"bufio"
	"io"
	"os"
)

type History struct {
	lines []string
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
				history.lines = append(history.lines, tmpline)
				tmpline = ""
			}
			tmpline += string(c)
		}
	}
	return nil
}

func (history *History) GetLines() []string {
	return history.lines
}
