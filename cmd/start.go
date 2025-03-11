package cmd

import (
	"fmt"
	"time"

	"github.com/temryakov/pomodoro/constants"
	"github.com/temryakov/pomodoro/entities"
	"github.com/temryakov/pomodoro/utils"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   constants.StartUse,
	Short: constants.StartShort,
	Long:  constants.StartLong,
	Run:   RunPomodoro,
}

func RunPomodoro(cmd *cobra.Command, args []string) {

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

	//TODO: Put default time into config
	startCmd.Flags().IntP("duration", "d", 25, "Time duration of pomodoro")
}
