package utils

// func SetTimerWithContext(c context.Context, duration time.Duration) time.Duration {
// 	var (
// 		ctx, cancel = context.WithCancel(c)
// 		start       = time.Now() // Start count spent time
// 	)
// 	defer cancel()
// 	go func() {
// 		SelectOption(ctx)
// 		cancel()
// 	}()
// 	SetTimer(ctx, duration)
// 	return time.Since(start)
// }

// func SetTimer(ctx context.Context, duration time.Duration) {
// 	start_time := time.Now()
// 	end_time := start_time.Add(duration)
// 	m := big.NewInt(int64(time.Until(end_time).Minutes()))
// 	s := big.NewInt(int64(time.Until(end_time).Seconds()))
// 	s = s.Mod(s, big.NewInt(60))
// 	m = m.Mod(m, big.NewInt(60))

// 	fmt.Print("\n")
// 	fmt.Printf("\r\t⏳ %2d minutes %2d seconds", m, s)

// 	i := int64(1)
// 	ticker1 := time.NewTicker(1 * time.Second)
// 	for range ticker1.C {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			i++
// 			m = big.NewInt(int64(time.Until(end_time).Minutes()))
// 			s = big.NewInt(int64(time.Until(end_time).Seconds()))
// 			s = s.Mod(s, big.NewInt(60))
// 			m = m.Mod(m, big.NewInt(60))

// 			fmt.Printf("\r\t⏳ %2d minutes %2d seconds", m, s)
// 			if i > int64(duration/time.Second) {
// 				ticker1.Stop()
// 				return
// 			}
// 		}
// 	}
// }
