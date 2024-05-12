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
	m := int(spent.Minutes())
	s := int(spent.Seconds()) - m*60

	fmt.Printf("\r\t⌛️ Total spent time: %2d minutes %2d seconds\n", m, s)
	fmt.Println(pb.FinishDescription())
	pb.Sound()
}
