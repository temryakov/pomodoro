package cmd

import (
	"pomodoro/entities"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start new break",
	Long:  "A longer description",
	Run:   RunBreak,
}

func RunBreak(cmd *cobra.Command, args []string) {
	//TODO: Put configtime into config

	timeconfig := 5
	duration := time.Duration(timeconfig) * time.Minute
	selectCh := make(chan int)
	br := entities.NewBreak(duration)

	entities.Start(br, selectCh)
	utils.SetTimerWithSelect(duration, selectCh)
	entities.Finish(br)
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
