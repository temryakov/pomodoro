package utils

import (
	"fmt"
	"pomodoro/domain"
)

func Finish(status domain.IStatus) {
	//TODO: Make "spent time.Duration" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %v minutes\n", 25)
	fmt.Println(status.FinishDescription())
	status.Sound()
}
