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
		start       = time.Now() // Start count spent time
	)
	defer cancel()
	go func() {
		SelectOption(ctx)
		cancel()
	}()
	SetTimer(ctx, duration)
	return time.Since(start)
}

func SetTimer(ctx context.Context, duration time.Duration) {
	remaining := time.Now().Add(duration)

	// Retrieve the current remaining time, accounting for delays in execution related to
	//
	tr := getTimeRemaining(remaining)
	fmt.Printf(constants.Countdown, tr.minutes, tr.seconds)

	for range time.NewTicker(1 * time.Second).C {
		select {
		case <-ctx.Done():
			return
		default:
			tr := getTimeRemaining(remaining)
			if tr.total <= 0 {
				return
			}
			fmt.Printf(constants.Countdown, tr.minutes, tr.seconds)
		}
	}
}

func getTimeRemaining(t time.Time) countdown {
	current := time.Now()
	diff := t.Sub(current)

	total := int(diff.Seconds())

	return countdown{
		total:   total,
		minutes: int(total/60) % 60,
		seconds: int(total % 60),
	}
}
