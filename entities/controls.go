package entities

import (
	"fmt"
	"pomodoro/domain"
	"pomodoro/utils"
)

func Start(pb domain.PomodoroBreaker, selectCh chan int) {
	fmt.Print(pb.StartDescription())
	go func() {
		result := utils.SelectOption()
		selectCh <- result
		close(selectCh)
	}()
}

func Finish(pb domain.PomodoroBreaker) {
	//TODO: Make "spent time.Duration" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %v minutes\n", 25)
	fmt.Println(pb.FinishDescription())
	pb.Sound()
}
