package pomodoro

import (
	"fmt"
	"log"
	"os/exec"
)

func Sound(state string) {
	var sound string
	if state == "break" {
		sound = "Blow"
	}
	if state == "pomodoro" {
		sound = "Submarine"
	}
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
