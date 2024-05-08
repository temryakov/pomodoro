package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"pomodoro/utils"
	"time"

	"github.com/spf13/cobra"
)

type Pomodoro struct {
	StartDescription  string
	FinishDescription string
	Duration          int //TODO: Fix Duration type to time.Duration
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start new pomodoro",
	Long:  "A longer description",
	Run:   RunPomodoro,
}

func NewPomodoro(Duration int) Pomodoro {
	return Pomodoro{
		StartDescription:  fmt.Sprintf("pomodoro: 🍅 Pomodoro has been started! it will take %v minutes. Don't forget to take a break.\n(In order to finish pomodoro, press key 1)\n", Duration),
		FinishDescription: "\npomodoro: 🍅 Finished! Print 'pomodoro break' to take a break.",
		Duration:          Duration,
	}
}

func (p Pomodoro) Start(selectCh chan int) {
	fmt.Print(p.StartDescription)
	go func() {
		result := utils.SelectOption()
		selectCh <- result
		close(selectCh)
	}()
}

func (p Pomodoro) Finish() {
	// TODO: Make "spent int" argument to pass spent time

	fmt.Printf("\r\t⌛️ Total spent time: %v minutes\n", 25)
	fmt.Println(p.FinishDescription)
	p.Sound()
}

func (p Pomodoro) Sound() {
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", "Submarine"))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func RunPomodoro(cmd *cobra.Command, args []string) {
	timer := 25
	pomodoro := NewPomodoro(timer)

	selectCh := make(chan int)
	pomodoro.Start(selectCh)

	duration := time.Duration(time.Minute * 25)
	utils.SetTimerWithSelect(duration, selectCh)
	pomodoro.Finish()
}

func init() {
	rootCmd.AddCommand(startCmd)
}
