package utils

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func SetTimerWithContext(c context.Context, duration time.Duration) time.Duration /*TODO: Implement return spent time value */ {
	var (
		ctx, cancel = context.WithCancel(c)
		wg          = sync.WaitGroup{}
		start       = time.Now()
	)
	defer cancel()
	wg.Add(1)
	go func() {
		SetTimer(ctx, duration)
		wg.Done()
		cancel()
	}()
	go func() {
		SelectOption(ctx)
		cancel()
	}()
	wg.Wait()
	return time.Since(start)
}

func SetTimer(ctx context.Context, duration time.Duration) {
	start_time := time.Now()
	end_time := start_time.Add(duration)
	m := big.NewInt(int64(time.Until(end_time).Minutes()))
	s := big.NewInt(int64(time.Until(end_time).Seconds()))
	s = s.Mod(s, big.NewInt(60))
	m = m.Mod(m, big.NewInt(60))

	fmt.Print("\n")
	fmt.Printf("\r\t⏳ %2d minutes %2d seconds", m, s)

	i := int64(1)
	ticker1 := time.NewTicker(1 * time.Second)
	for range ticker1.C {
		select {
		case <-ctx.Done():
			return
		default:
			i++
			m = big.NewInt(int64(time.Until(end_time).Minutes()))
			s = big.NewInt(int64(time.Until(end_time).Seconds()))
			s = s.Mod(s, big.NewInt(60))
			m = m.Mod(m, big.NewInt(60))

			fmt.Printf("\r\t⏳ %2d minutes %2d seconds", m, s)
			if i > int64(duration/time.Second) {
				ticker1.Stop()
				return
			}
		}
	}
}
