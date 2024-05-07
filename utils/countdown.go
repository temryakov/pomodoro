package utils

import (
	"fmt"
	"math/big"
	"time"
)

func SetTimerWithSelect(duration time.Duration, selectCh chan int) {
	done := make(chan bool)
	go func() {
		SetTimer(duration)
		done <- true
	}()
	select {
	case m := <-selectCh:
		if m == StatusFinish {
			return
		}
	case <-done:
		return
	}
}

func SetTimer(duration time.Duration) {
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
