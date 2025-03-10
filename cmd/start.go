package cmd

import (
	"fmt"
	"pomodoro/constants"
	"pomodoro/entities"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   constants.StartUse,
	Short: constants.StartShort,
	Long:  constants.StartLong,
	Run:   RunPomodoro,
}

func RunPomodoro(cmd *cobra.Command, args []string) {
	//TODO: Put configtime into config

	timeconfig, _ := cmd.Flags().GetInt("duration")

	duration := time.Duration(timeconfig) * time.Minute
	p := entities.NewPomodoro(duration)

	fmt.Print(p.StartDescription())
	spent := utils.SetTimerWithContext(duration)

	fmt.Print(constants.ErasingString)
	fmt.Print(utils.GetTimeSpentString(spent))
	fmt.Println(p.FinishDescription())
	p.Sound()
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("duration", "d", 25, "Time duration of pomodoro")
}
