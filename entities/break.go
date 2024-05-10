package entities

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Break struct {
	startDescription  string
	finishDescription string
	Duration          time.Duration
}

func NewBreak(Duration time.Duration) Break {
	return Break{
		startDescription:  fmt.Sprintf("pomodoro: ☕️ Break has been started! it will take %v minutes. Have a good time!\n(In order to finish break, press key 1)\n", Duration.Minutes()),
		finishDescription: "\npomodoro: ☕️ It's time to get work! Print 'pomodoro start' to start new pomodoro.",
		Duration:          Duration,
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
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
