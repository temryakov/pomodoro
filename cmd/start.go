package cmd

import (
	"context"
	"pomodoro/entities"
	"pomodoro/service"
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

	timeconfig, _ := cmd.Flags().GetInt("duration")

	duration := time.Duration(timeconfig) * time.Minute
	ctx := context.Background()
	p := entities.NewPomodoro(duration)

	service.Start(p)
	spent := utils.SetTimerWithContext(ctx, duration)
	service.Finish(p, spent)
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("duration", "d", 25, "Time duration of pomodoro")
}
