package entities

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/constants"
	"time"
)

type Pomodoro struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
}

func NewPomodoro(Duration time.Duration) Pomodoro {
	return Pomodoro{
		startDescription:  fmt.Sprintf(constants.PomodoroStartDesc, Duration.Minutes()),
		finishDescription: constants.PomodoroFinishDesc,
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
	cmd := exec.Command("afplay", fmt.Sprintf("./sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
