package cmd

import (
	"context"
	"pomodoro/entities"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start new pomodoro",
	Long:  "A longer description",
	Run:   RunPomodoro,
}

func RunPomodoro(cmd *cobra.Command, args []string) {
	//TODO: Put configtime into config

	timeconfig := 25
	duration := time.Duration(timeconfig) * time.Minute
	ctx := context.Background()
	p := entities.NewPomodoro(duration)

	entities.Start(p)
	spent := utils.SetTimerWithContext(ctx, duration)
	entities.Finish(p, spent)
}

func init() {
	rootCmd.AddCommand(startCmd)
}
