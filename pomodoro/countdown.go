package pomodoro

import (
	"fmt"
	"math/big"
	"time"
)

func SetTimer(duration time.Duration) {
	start_time := time.Now()
	end_time := start_time.Add(duration)
	m := big.NewInt(int64(time.Until(end_time).Minutes()))
	s := big.NewInt(int64(time.Until(end_time).Seconds()))
	s = s.Mod(s, big.NewInt(60))
	m = m.Mod(m, big.NewInt(60))

	fmt.Print("\n")
	fmt.Printf("\r\t⏳ %2d minutes %2d seconds", m, s)
	fmt.Print("\r")

	i := int64(1)
	ticker1 := time.NewTicker(1 * time.Second)
	for range ticker1.C {
		i++
		m = big.NewInt(int64(time.Until(end_time).Minutes()))
		s = big.NewInt(int64(time.Until(end_time).Seconds()))
		s = s.Mod(s, big.NewInt(60))
		m = m.Mod(m, big.NewInt(60))

		fmt.Printf("\r\t⏳ %2d minutes %2d seconds", m, s)
		fmt.Print("\r")

		if i > int64(duration/time.Second) {
			ticker1.Stop()
			break
		}
	}
	fmt.Printf("\r\t⏳ Total spent time: 25 minutes\n")
}
