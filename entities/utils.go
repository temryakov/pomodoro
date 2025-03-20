package entities

import (
	"errors"
	"fmt"
	"time"
)

func buildResult(spent time.Duration) (string, error) {
	var (
		m  = int(spent.Minutes())
		s  = int(spent.Seconds()) - m*60
		ms string
		ss string
	)
	if m > 0 {
		ms = fmt.Sprintf("%2d minutes ", m)
	}
	if s > 0 {
		ss = fmt.Sprintf("%2d seconds ", s)
	}

	if ms == "" && ss == "" {
		return "", errors.New("empty duration")
	}
	return ms + ss, nil
}
