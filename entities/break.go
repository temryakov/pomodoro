package entities

import (
	"fmt"
	"time"

	"github.com/temryakov/pomodoro/app"
	"github.com/temryakov/pomodoro/constants"
	"github.com/temryakov/pomodoro/domain"
)

type Break struct {
	startDescription  string
	finishDescription string
	duration          time.Duration
	sound             string
	Repository        domain.Repository
}

func NewBreak(duration time.Duration, repository domain.Repository) Break {
	return Break{
		startDescription:  fmt.Sprintf(constants.BreakStartDesc, duration.Minutes(), app.StatusPause, app.StatusFinish),
		finishDescription: constants.BreakFinishDesc,
		duration:          duration,
		sound:             constants.BreakSound,
		Repository:        repository,
	}
}

func (b Break) StartDescription() string {
	return b.startDescription
}

func (b Break) FinishDescription() string {
	return b.finishDescription
}

func (b Break) Sound() {
	app.ExecSound(b.sound)
}

func (b Break) SaveHistory(duration time.Duration) {
	if duration.Minutes() < 1 {
		return
	}
	err := b.Repository.Post(duration, constants.BreakRecord)
	if err != nil {
		panic(err)
	}
}

func (b Break) GetLast() {
	res, err := b.Repository.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
