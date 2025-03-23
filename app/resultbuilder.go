package app

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/temryakov/pomodoro/domain"
)

const (
	tableIsEmpty    = "\npomodoro: 📭 table is empty"
	timeSpentHeader = "\r\t⌛️ Total spent time: "
)

func GetTimeSpentString(spent time.Duration) string {
	var (
		m      = int(spent.Minutes())
		s      = int(spent.Seconds()) - m*60
		header = timeSpentHeader
		ms     string
		ss     string
	)
	if m > 0 {
		ms = fmt.Sprintf("%2d minutes ", m)
	}
	if s > 0 {
		ss = fmt.Sprintf("%2d seconds ", s)
	}

	if ms == "" && ss == "" {
		return "\r\t⌛️ Done! \n"
	}
	return header + ms + ss + "\n"
}

func GetHistoryList(list []domain.History) {
	if len(list) == 0 {
		fmt.Println(tableIsEmpty)
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	// Table header
	fmt.Fprintln(w, "\npomodoro: 🍅 History list of pomodoros and breaks: ")
	// fmt.Fprintln(w, "\nName\tDuration\tEnd Date")
	fmt.Fprintln(w, "\t\t")

	var prevDate string

	for _, h := range list {
		curr := h.Date
		currentDate := curr.Format("2 Jan 2006")
		// If previous day is not equal than current, make separator
		if prevDate != currentDate {
			fmt.Fprintf(w, "%s\tDuration --\tEnd Date --\n", currentDate)
			fmt.Fprintln(w, "\t\t")
			prevDate = currentDate
		}
		// Main string with data
		date := curr.Format("15:04")
		fmt.Fprintf(w, "%s\t%s\t%s\n", h.Name, GetPomodoroResult(h.Duration), date)
	}
	w.Flush()
}

func GetPomodoroResult(spent time.Duration) string {
	var (
		m  = int(spent.Minutes())
		s  = int(spent.Seconds()) - m*60
		ms string
		ss string
	)
	if m > 0 {
		ms = fmt.Sprintf("%2d min ", m)
	}
	if s > 0 {
		ss = fmt.Sprintf("%2d sec ", s)
	}
	return ms + ss
}
