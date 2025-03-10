package utils

import (
	"fmt"
	"pomodoro/constants"

	"github.com/eiannone/keyboard"
)

var (
	StatusPause  int = 0
	StatusFinish int = 1
)

func (t *Timer) SelectOption() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
		select {
		case <-t.Finish:
			return
		default:
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			if int(char-'0') == StatusFinish {
				fmt.Print(constants.ErasingString)
				fmt.Print(constants.FinishingProcess)
				t.Pause <- struct{}{}
				t.Finish <- struct{}{}
				return
			}
			if int(char-'0') == StatusPause {
				fmt.Print(constants.ErasingString)
				fmt.Print(constants.PausingProcess)
				t.Pause <- struct{}{}
				continue
			}
		}
	}
}
