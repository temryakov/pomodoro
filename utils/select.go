package utils

import (
	"context"
	"fmt"

	"github.com/eiannone/keyboard"
)

var (
	StatusPause  int = 0
	StatusFinish int = 1
)

func SelectOption(ctx context.Context) {
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
				fmt.Print("\r\t⏳ Finishing...         ")
				return
			}
		}
	}
}
