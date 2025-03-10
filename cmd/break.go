package cmd

import (
	"fmt"
	"pomodoro/constants"
	"pomodoro/entities"
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
	br := entities.NewBreak(duration)

	fmt.Print(br.StartDescription())
	spent := utils.SetTimerWithContext(duration)

	fmt.Print(constants.ErasingString)
	fmt.Print(utils.GetTimeSpentString(spent))
	fmt.Println(br.FinishDescription())
	br.Sound()
}

func init() {
	rootCmd.AddCommand(breakCmd)
	breakCmd.Flags().IntP("duration", "d", 5, "Time duration of break")
}
