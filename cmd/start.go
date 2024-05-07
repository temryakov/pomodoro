package cmd

import (
	"pomodoro/domain"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

var pomodoro = domain.NewPomodoro(25)

var startCmd = &cobra.Command{
	Use:   pomodoro.Use,
	Short: pomodoro.ShortDescription,
	Long:  pomodoro.LongDescription,
	Run:   RunPomodoro,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func RunPomodoro(cmd *cobra.Command, args []string) {
	// timer := 25
	// pomodoro := domain.NewPomodoro(timer)

	selectCh := make(chan int)
	pomodoro.Start(selectCh)

	duration := time.Duration(time.Minute * 25)
	utils.SetTimerWithSelect(duration, selectCh)
	pomodoro.Finish()
}
