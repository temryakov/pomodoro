package cmd

import (
	"context"
	"pomodoro/constants"
	"pomodoro/entities"
	"pomodoro/service"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

var breakCmd = &cobra.Command{
	Use:   constants.BreakUse,
	Short: constants.BreakShort,
	Long:  constants.BreakLong,
	Run:   RunBreak,
}

func RunBreak(cmd *cobra.Command, args []string) {
	//TODO: Put configtime into config

	timeconfig, _ := cmd.Flags().GetInt("duration")

	duration := time.Duration(timeconfig) * time.Minute
	ctx := context.Background()
	br := entities.NewBreak(duration)

	service.Start(br)
	spent := utils.SetTimerWithContext(ctx, duration)
	service.Finish(br, spent)
}

func init() {
	rootCmd.AddCommand(breakCmd)
	breakCmd.Flags().IntP("duration", "d", 5, "Time duration of break")
}
