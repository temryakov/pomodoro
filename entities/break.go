package entities

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/constants"
	"time"
)

type Break struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
}

func NewBreak(Duration time.Duration) Break {
	return Break{
		startDescription:  fmt.Sprintf(constants.BreakStartDesc, Duration.Minutes()),
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
