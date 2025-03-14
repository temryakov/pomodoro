package cmd

import (
	"fmt"
	"time"

	"github.com/temryakov/pomodoro/app"
	"github.com/temryakov/pomodoro/constants"
	"github.com/temryakov/pomodoro/entities"
	"github.com/temryakov/pomodoro/repository"

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

	r := repository.NewRepository(app.InitDB())

	duration := time.Duration(timeconfig) * time.Minute
	p := entities.NewPomodoro(duration, r)

	defer p.Repository.Close()

	fmt.Print(p.StartDescription())
	spent := app.SetTimerWithContext(duration)

	p.SaveHistory(spent)

	fmt.Print(constants.ErasingString)
	fmt.Print(app.GetTimeSpentString(spent))
	fmt.Println(p.FinishDescription())
	p.Sound()
}

func init() {
	rootCmd.AddCommand(startCmd)

	//TODO: Put default time into config
	startCmd.Flags().IntP("duration", "d", 25, "Time duration of pomodoro")
}
