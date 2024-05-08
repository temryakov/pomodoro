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
	Duration          int //TODO: Fix Duration type to time.Duration
}

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start new break",
	Long:  "A longer description",
	Run:   RunBreak,
}

func NewBreak(Duration int) Break {
	return Break{
		StartDescription:  fmt.Sprintf("pomodoro: ☕️ Break has been started! it will take %v minutes. Have a good time!\n(In order to finish break, press key 1)\n", Duration),
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

func (p Break) Finish() {
	// TODO: Make "spent int" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %v minutes\n", 5)
	fmt.Println(p.FinishDescription)
	p.Sound()
}

func (b Break) Sound() {
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", "Blow"))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func RunBreak(cmd *cobra.Command, args []string) {
	timer := 5
	br := NewBreak(timer)
	selectCh := make(chan int)
	br.Start(selectCh)

	duration := time.Duration(time.Minute * 5)
	utils.SetTimerWithSelect(duration, selectCh)
	br.Finish()
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
