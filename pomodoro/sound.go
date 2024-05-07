package pomodoro

import (
	"log"
	"os/exec"
)

func Sound() {
	cmd := exec.Command("afplay", "/System/Library/Sounds/Submarine.aiff")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
