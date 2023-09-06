package history

import (
	"errors"
	"time"

	"github.com/mrusme/shell-time/history/zsh"
)

type History interface {
	Initialize(path string) error
	GetNumberOfLines() int64
	GetLines() []*string
	GetLine(lineIdx int64) (time.Time, string, string, error)
}

func New(historyFormat string, historyFile string) (History, error) {
	var hist History

	switch historyFormat {
	case "zsh":
		hist = new(zsh.History)
	default:
		return nil, errors.New("No such history")
	}

	err := hist.Initialize(historyFile)
	if err != nil {
		return nil, err
	}

	return hist, nil
}
