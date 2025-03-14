package app

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func ExecSound(name string) {
	if runtime.GOOS != "darwin" {
		return
	}
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", name))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
