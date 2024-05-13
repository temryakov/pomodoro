package utils

import (
	"fmt"
	"time"
)

func GetTimeSpentString(spent time.Duration) string {
	var (
		m      = int(spent.Minutes())
		s      = int(spent.Seconds()) - m*60
		header = "\r\t⌛️ Total spent time: "
		ms     string
		ss     string
	)
	if m > 0 {
		ms = fmt.Sprintf("%2d minutes ", m)
	}
	if s > 0 {
		ss = fmt.Sprintf("%2d seconds ", s)
	}

	return header + ms + ss + "\n"
}
