package entities

import (
	"fmt"
	"time"

	"github.com/temryakov/pomodoro/app"
	"github.com/temryakov/pomodoro/constants"
	"github.com/temryakov/pomodoro/domain"
)

type Pomodoro struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
	sound             string
	Repository        domain.Repository
}

func NewPomodoro(Duration time.Duration, repository domain.Repository) Pomodoro {
	return Pomodoro{
		startDescription:  fmt.Sprintf(constants.PomodoroStartDesc, Duration.Minutes(), app.StatusPause, app.StatusFinish),
		finishDescription: constants.PomodoroFinishDesc,
		duration:          Duration,
		sound:             "Submarine", // TODO: put into constants
		Repository:        repository,
	}
}

func (p Pomodoro) StartDescription() string {
	return p.startDescription
}

func (p Pomodoro) FinishDescription() string {
	return p.finishDescription
}

func (p Pomodoro) Sound() {
	app.ExecSound(p.sound)
}

func (p Pomodoro) SaveHistory(duration time.Duration) {

	res, err := buildResult(duration)

	// If got error, just do nothing
	if err != nil {
		return
	}
	err = p.Repository.Post(res, domain.PomodoroRecord)
	if err != nil {
		panic(err)
	}
}

func (p Pomodoro) GetLast() {
	res, err := p.Repository.Get()
	if err != nil {
		panic(err)
	}
	app.GetHistoryList(res)
}
