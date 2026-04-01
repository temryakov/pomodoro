package app

import (
	"fmt"
	"time"

	"github.com/temryakov/pomodoro/constants"
)

type countdown struct {
	total, minutes, seconds int
}

type Timer struct {
	Finish   chan struct{}
	Pause    chan struct{}
	Duration time.Duration
}

func SetTimerWithContext(duration time.Duration) time.Duration {
	timer := &Timer{
		Finish:   make(chan struct{}),
		Pause:    make(chan struct{}),
		Duration: duration,
	}

	go func() {
		timer.SelectOption()
		timer.Finish <- struct{}{}
	}()
	return duration - timer.SetTimer()
}

func (t *Timer) SetTimer() time.Duration {
	// Is pause right now?
	var pause bool
	var pausedAt time.Time

	start := time.Now()
	originalDuration := t.Duration

	for range time.NewTicker(100 * time.Millisecond).C {
		select {
		case <-t.Pause:
			pause = true
			pausedAt = time.Now()
		case <-t.Finish:
			return t.Duration
		default:
			if pause {
				<-t.Pause
				// Shift start forward by the time we were paused
				start = start.Add(time.Since(pausedAt))
				pause = false
			}
			elapsed := time.Since(start)
			remaining := originalDuration - elapsed
			t.Duration = remaining
			tr := getTimeRemaining(remaining)
			if tr.total <= 0 {
				t.Duration = 0
				return 0
			}
			fmt.Printf(constants.Countdown, tr.minutes, tr.seconds)
		}
	}
	return t.Duration
}

func getTimeRemaining(t time.Duration) countdown {

	total := int(t.Seconds())

	return countdown{
		total:   total,
		minutes: int(total/60) % 60,
		seconds: int(total % 60),
	}
}
