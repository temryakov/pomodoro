package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

type Break struct {
	StartDescription  string
	FinishDescription string
	Duration          time.Duration
}

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start new break",
	Long:  "A longer description",
	Run:   RunBreak,
}

func NewBreak(Duration time.Duration) Break {
	return Break{
		StartDescription:  fmt.Sprintf("pomodoro: ☕️ Break has been started! it will take %v minutes. Have a good time!\n(In order to finish break, press key 1)\n", Duration.Minutes()),
		FinishDescription: "\npomodoro: ☕️ It's time to get work! Print 'pomodoro start' to start new pomodoro.",
		Duration:          Duration,
	}
}

func (b Break) Start(selectCh chan int) {
	fmt.Printf(b.StartDescription)
	go func() {
		result := utils.SelectOption()
		selectCh <- result
		close(selectCh)
	}()
}

func (b Break) GetFinishDescription() string {
	return b.FinishDescription
}

func (b Break) Sound() {
	sound := "Blow" //TODO: Put into config
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func RunBreak(cmd *cobra.Command, args []string) {
	//TODO: Put configtime into config
	timeconfig := 5

	duration := time.Duration(timeconfig) * time.Minute
	br := NewBreak(duration)
	selectCh := make(chan int)
	br.Start(selectCh)

	utils.SetTimerWithSelect(duration, selectCh)
	utils.Finish(br)
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
