package utils

import "github.com/eiannone/keyboard"

var (
	StatusPause  int = 0
	StatusFinish int = 1
)

func SelectOption() int {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		switch int(char - '0') {
		case StatusFinish:
			return StatusFinish
		}
	}
}
