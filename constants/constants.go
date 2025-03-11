package constants

const (
	RootUse   = "pomodoro"
	RootShort = "🍅 Pomodoro CLI for your work and study"
	RootLong  = `🍅 Pomodoro CLI is a productivity cli application designed 
for your work and study. Stay focused and finish tasks effectively. 
Take more time to your life.`
)

const (
	StartUse   = "start"
	StartShort = "Start new pomodoro"
	StartLong  = "A longer description"
)

const (
	BreakUse   = "break"
	BreakShort = "Start new break"
	BreakLong  = "A longer description"
)

const (
	PomodoroStartDesc  = "pomodoro: 🍅 Pomodoro has been started! it will take %v minutes. Don't forget to take a break.\n(In order to pause pomodoro, press key %d or press %d to finish)\n\n"
	PomodoroFinishDesc = "\npomodoro: 🍅 Finished! Print 'pomodoro break' to take a break."

	BreakStartDesc  = "pomodoro: ☕️ Break has been started! it will take %v minutes. Have a good time!\n(In order to pause break, press key %d or press %d to finish)\n\n"
	BreakFinishDesc = "\npomodoro: ☕️ It's time to get work! Print 'pomodoro start' to start new pomodoro."

	Countdown        = "\r\t⏳ %2d minutes %2d seconds"
	FinishingProcess = "\r\t⏳ Finishing..."
	PausingProcess   = "\r\t💤 Pause."

	ErasingString = "\r\t                                    "
)
