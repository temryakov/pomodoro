package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/temryakov/pomodoro/constants"
	"github.com/temryakov/pomodoro/domain"
	"github.com/temryakov/pomodoro/utils"
)

type Pomodoro struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
	sound             string
	Repository        domain.Repository
}

func NewPomodoro(Duration time.Duration /*, repository domain.Repository */) Pomodoro {
	return Pomodoro{
		startDescription:  fmt.Sprintf(constants.PomodoroStartDesc, Duration.Minutes(), utils.StatusPause, utils.StatusFinish),
		finishDescription: constants.PomodoroFinishDesc,
		duration:          Duration,
		sound:             "Submarine", // TODO: put into constants
		// Repository:        repository,
	}
}

func (p Pomodoro) StartDescription() string {
	return p.startDescription
}

func (p Pomodoro) FinishDescription() string {
	return p.finishDescription
}

func (p Pomodoro) Sound() {
	utils.ExecSound(p.sound)
}

func (p Pomodoro) SaveHistory(duration time.Duration) {

	res, err := buildResult(duration)
	if err != nil {
		return
	}
	p.Repository.Post(res)
}

func (p Pomodoro) GetLast() {
	res, err := p.Repository.Get()
	if err != nil {
		return
	}
	fmt.Println(res)
}

func buildResult(spent time.Duration) (string, error) {
	var (
		m  = int(spent.Minutes())
		s  = int(spent.Seconds()) - m*60
		ms string
		ss string
	)
	if m > 0 {
		ms = fmt.Sprintf("%2d minutes ", m)
	}
	if s > 0 {
		ss = fmt.Sprintf("%2d seconds ", s)
	}

	if ms == "" && ss == "" {
		return "", errors.New("empty duration")
	}
	return ms + ss, nil
}
