package app

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/temryakov/pomodoro/domain"
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

	if ms == "" && ss == "" {
		return "\r\t⌛️ Done! \n"
	}
	return header + ms + ss + "\n"
}

func GetHistoryList(list []domain.History) {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)

	fmt.Fprintln(w, "Name\tDuration\tEnd Date\t")
	fmt.Fprintln(w, "----\t--------\t---------\t")

	for _, h := range list {
		date := h.Date.Format("2 Jan, 15:04")
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", h.Name, h.Duration, date)
	}

	w.Flush()
}
