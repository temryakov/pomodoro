package entities

import (
	"fmt"
	"pomodoro/domain"
	"time"
)

func Start(pb domain.PomodoroBreaker) {
	fmt.Print(pb.StartDescription())
}

func Finish(pb domain.PomodoroBreaker, spent time.Duration) {
	//TODO: Make "spent time.Duration" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %2d minutes\n", int(spent.Minutes()))
	fmt.Println(pb.FinishDescription())
	pb.Sound()
}
