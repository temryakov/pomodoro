package cmd

import (
	"pomodoro/domain"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

var br = domain.NewBreak(5)

var breakCmd = &cobra.Command{
	Use:   br.Use,
	Short: br.ShortDescription,
	Long:  br.LongDescription,
	Run:   RunBreak,
}

func init() {
	rootCmd.AddCommand(breakCmd)
}

func RunBreak(cmd *cobra.Command, args []string) {

	selectCh := make(chan int)
	br.Start(selectCh)

	duration := time.Duration(time.Minute * 5)
	utils.SetTimerWithSelect(duration, selectCh)
	br.Finish()
}
