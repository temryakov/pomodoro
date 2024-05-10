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
	Duration          time.Duration
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start new pomodoro",
	Long:  "A longer description",
	Run:   RunPomodoro,
}

func NewPomodoro(Duration time.Duration) Pomodoro {
	return Pomodoro{
		StartDescription:  fmt.Sprintf("pomodoro: 🍅 Pomodoro has been started! it will take %v minutes. Don't forget to take a break.\n(In order to finish pomodoro, press key 1)\n", Duration.Minutes()),
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

func (p Pomodoro) GetFinishDescription() string {
	return p.FinishDescription
}

func (p Pomodoro) Sound() {
	sound := "Submarine" //TODO: Put into config
	cmd := exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%v.aiff", sound))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func RunPomodoro(cmd *cobra.Command, args []string) {
	//TODO: Put configtime into config
	configtime := 25

	duration := time.Duration(configtime) * time.Minute
	pomodoro := NewPomodoro(duration)

	selectCh := make(chan int)
	pomodoro.Start(selectCh)

	utils.SetTimerWithSelect(duration, selectCh)
	utils.Finish(pomodoro)
}

func init() {
	rootCmd.AddCommand(startCmd)
}
