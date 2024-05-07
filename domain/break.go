package domain

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/utils"
)

type Break struct {
	Use               string
	ShortDescription  string
	LongDescription   string
	StartDescription  string
	FinishDescription string
	Duration          int
	HasFinish         bool
	HasPause          bool
}

func NewBreak(Duration int) Break {
	return Break{
		Use:               "break",
		ShortDescription:  "Start new break",
		LongDescription:   "A longer description",
		StartDescription:  fmt.Sprintf("pomodoro: ☕️ Break has been started! it will take %v minutes. Have a good time!\n(In order to finish break, press key 1)\n", Duration),
		FinishDescription: "\npomodoro: ☕️ It's time to get work! Print 'pomodoro start' to start new pomodoro.",
		Duration:          Duration,
		HasFinish:         true,
		HasPause:          false,
	}
}

func (b Break) Start(selectCh chan int) {
	fmt.Printf(b.StartDescription)
	go func() {
		result := utils.SelectOption()
		selectCh <- result
		close(selectCh)
	}()
}

func (p Break) Finish() {
	// TODO: Make "spent int" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %v minutes\n", 5)
	fmt.Println(p.FinishDescription)
	p.Sound()
}

func (b Break) Sound() {
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", "Blow"))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
