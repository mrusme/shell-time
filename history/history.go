package history

import (
	"errors"

	"github.com/mrusme/shell-time/history/zsh"
)

type History interface {
	Initialize(path string) error
	GetLines() []*string
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
