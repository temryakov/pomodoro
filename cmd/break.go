/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"pomodoro/pomodoro"
	"time"

	"github.com/spf13/cobra"
)

// breakCmd represents the break command
var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run:   RunBreak,
}

func init() {
	rootCmd.AddCommand(breakCmd)
}

func RunBreak(cmd *cobra.Command, args []string) {
	timer := 5
	fmt.Printf("pomodoro: ☕️ Break has been started! it will take %v minutes. Have a good time!\n(In order to finish break, press key 1)\n", timer)
	selectCh := make(chan int)
	go func() {
		result := SelectOption()
		selectCh <- result
		close(selectCh)
	}()
	duration := time.Duration(time.Minute * 5)
	pomodoro.SetTimerWithSelect(duration, selectCh)
	fmt.Printf("\r\t⏳ Total spent time: %v minutes\n", timer)
	log.Println("☕️ It's time to get work! Print 'pomodoro start' to start new pomodoro.")
	pomodoro.Sound("break")
}
