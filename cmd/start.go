/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"pomodoro/pomodoro"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start new pomodoro",
	Long:  `A longer description`,
	Run:   RunPomodoro,
}

const (
	statusPause  int = 2
	statusFinish int = 1
)

func init() {
	rootCmd.AddCommand(startCmd)
}

func RunPomodoro(cmd *cobra.Command, args []string) {
	timer := 25
	fmt.Printf("pomodoro: 🍅 Pomodoro has been started! it will take %v minutes. Don't forget to take a break.\n(In order to finish pomodoro, press key 1)\n", timer)
	selectCh := make(chan int)
	go func() {
		result := SelectOption()
		selectCh <- result
		close(selectCh)
	}()
	duration := time.Duration(time.Second * 5)
	pomodoro.SetTimerWithSelect(duration, selectCh)
	fmt.Printf("\r\t⏳ Total spent time: %v minutes\n", timer)
	log.Println("🍅 Finished! Print 'pomodoro break' to take a break.")
	pomodoro.Sound()
}

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
		switch char - '0' {
		case rune(statusFinish):
			return statusFinish
		}
	}
}
