package entities

import (
	"fmt"
	"time"

	"github.com/temryakov/pomodoro-cli/constants"
	"github.com/temryakov/pomodoro-cli/utils"
)

type Break struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
	sound             string
}

func NewBreak(Duration time.Duration) Break {
	return Break{
		startDescription:  fmt.Sprintf(constants.BreakStartDesc, Duration.Minutes(), utils.StatusPause, utils.StatusFinish),
		finishDescription: constants.BreakFinishDesc,
		duration:          Duration,
		sound:             "Blow", // TODO: put in constants
	}
}

func (b Break) StartDescription() string {
	return b.startDescription
}

func (b Break) FinishDescription() string {
	return b.finishDescription
}

func (b Break) Sound() {
	utils.ExecSound(b.sound)
}
