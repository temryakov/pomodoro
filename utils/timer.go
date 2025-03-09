package utils

import (
	"context"
	"fmt"
	"pomodoro/constants"
	"time"
)

type countdown struct {
	total, minutes, seconds int
}

func SetTimerWithContext(c context.Context, duration time.Duration) time.Duration {
	var (
		ctx, cancel = context.WithCancel(c)
		pauseCh     = make(chan struct{})
	)
	defer cancel()
	go func() {
		SelectOption(ctx, pauseCh)
		cancel()
	}()
	return duration - SetTimer(ctx, duration, pauseCh)
}

func SetTimer(ctx context.Context, duration time.Duration, pauseCh chan struct{}) time.Duration {
	var pause bool
	// Retrieve the current remaining time, accounting for delays in execution
	// related to non-instant work of Ticker.
	tr := getTimeRemaining(duration)
	fmt.Printf(constants.Countdown, tr.minutes, tr.seconds)

	for range time.NewTicker(1 * time.Second).C {
		select {
		case <-pauseCh:
			pause = true
			duration += time.Second
		case <-ctx.Done():
			return duration
		default:
			if pause {
				<-pauseCh
				pause = false
			}
			duration -= time.Second
			tr := getTimeRemaining(duration)
			if tr.total <= 0 {
				return duration
			}
			fmt.Printf(constants.Countdown, tr.minutes, tr.seconds)
		}
	}
	return duration
}

func getTimeRemaining(t time.Duration) countdown {

	total := int(t.Seconds())

	return countdown{
		total:   total,
		minutes: int(total/60) % 60,
		seconds: int(total % 60),
	}
}
