package domain

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/utils"
)

type Pomodoro struct {
	Use               string
	ShortDescription  string
	LongDescription   string
	StartDescription  string
	FinishDescription string
	Duration          int
	HasFinish         bool
	HasPause          bool
}

func NewPomodoro(Duration int) Pomodoro {
	return Pomodoro{
		Use:               "start",
		ShortDescription:  "Start new pomodoro",
		LongDescription:   "A longer description",
		StartDescription:  fmt.Sprintf("pomodoro: 🍅 Pomodoro has been started! it will take %v minutes. Don't forget to take a break.\n(In order to finish pomodoro, press key 1)\n", Duration),
		FinishDescription: "\npomodoro: 🍅 Finished! Print 'pomodoro break' to take a break.",
		Duration:          Duration,
		HasFinish:         true,
		HasPause:          false,
	}
}

func (p Pomodoro) Start(selectCh chan int) {
	fmt.Print(p.StartDescription)
	go func() {
		result := utils.SelectOption()
		selectCh <- result
		close(selectCh)
	}()
}

func (p Pomodoro) Finish() {
	// TODO: Make "spent int" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %v minutes\n", 25)
	fmt.Println(p.FinishDescription)
	p.Sound()
}

func (p Pomodoro) Sound() {
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", "Submarine"))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
