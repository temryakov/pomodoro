package entities

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/constants"
	"pomodoro/utils"
	"time"
)

type Break struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
}

func NewBreak(Duration time.Duration) Break {
	return Break{
		startDescription:  fmt.Sprintf(constants.BreakStartDesc, Duration.Minutes(), utils.StatusPause, utils.StatusFinish),
		finishDescription: constants.BreakFinishDesc,
		duration:          Duration,
	}
}

func (b Break) StartDescription() string {
	return b.startDescription
}

func (b Break) FinishDescription() string {
	return b.finishDescription
}

func (b Break) Sound() {
	sound := "Blow" //TODO: Put into config
	cmd := exec.Command("afplay", fmt.Sprintf("./sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
