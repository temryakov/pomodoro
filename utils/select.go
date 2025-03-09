package utils

import (
	"context"
	"fmt"
	"pomodoro/constants"

	"github.com/eiannone/keyboard"
)

var (
	StatusPause  int = 0
	StatusFinish int = 1
)

func SelectOption(ctx context.Context, pauseCh chan struct{}) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("SelectOption ctx.Done")
			return
		default:
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			if int(char-'0') == StatusFinish {
				fmt.Print(constants.ErasingString)
				fmt.Print(constants.FinishingProcess)
				pauseCh <- struct{}{}
				return
			}
			if int(char-'0') == StatusPause {
				fmt.Print(constants.ErasingString)
				fmt.Print(constants.PausingProcess)
				pauseCh <- struct{}{}
				continue
			}
		}
	}
}
