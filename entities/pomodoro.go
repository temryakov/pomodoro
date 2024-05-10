package entities

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Pomodoro struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
}

func NewPomodoro(Duration time.Duration) Pomodoro {
	return Pomodoro{
		startDescription:  fmt.Sprintf("pomodoro: 🍅 Pomodoro has been started! it will take %v minutes. Don't forget to take a break.\n(In order to finish pomodoro, press key 1)\n", Duration.Minutes()),
		finishDescription: "\npomodoro: 🍅 Finished! Print 'pomodoro break' to take a break.",
		duration:          Duration,
	}
}

func (p Pomodoro) StartDescription() string {
	return p.startDescription
}

func (p Pomodoro) FinishDescription() string {
	return p.finishDescription
}

func (p Pomodoro) Sound() {
	sound := "Submarine" //TODO: Put into config
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
