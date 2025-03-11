package entities

import (
	"fmt"
	"pomodoro/constants"
	"pomodoro/utils"
	"time"
)

type Pomodoro struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
	sound             string
}

func NewPomodoro(Duration time.Duration) Pomodoro {
	return Pomodoro{
		startDescription:  fmt.Sprintf(constants.PomodoroStartDesc, Duration.Minutes(), utils.StatusPause, utils.StatusFinish),
		finishDescription: constants.PomodoroFinishDesc,
		duration:          Duration,
		sound:             "Submarine", // TODO: put into constants
	}
}

func (p Pomodoro) StartDescription() string {
	return p.startDescription
}

func (p Pomodoro) FinishDescription() string {
	return p.finishDescription
}

func (p Pomodoro) Sound() {
	utils.ExecSound(p.sound)
}
