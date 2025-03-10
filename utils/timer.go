package utils

import (
	"fmt"
	"pomodoro/constants"
	"time"
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
	duration += time.Second

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

	for range time.NewTicker(1 * time.Millisecond).C {
		select {
		case <-t.Pause:
			pause = true
		case <-t.Finish:
			return t.Duration
		default:
			if pause {
				<-t.Pause
				pause = false
			}
			tr := getTimeRemaining(t.Duration)
			if tr.total <= 0 {
				return t.Duration
			}
			t.Duration -= time.Millisecond
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
