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
	StartLong  = `The pomodoro start command initiates a Pomodoro session, which is typically 
a focused work period lasting 25 minutes. This command is designed to help you concentrate on your tasks 
without distractions. When you run this command, the timer begins, and you can fully immerse yourself in your work.`
)

const (
	BreakUse   = "break"
	BreakShort = "Start new break"
	BreakLong  = `The pomodoro break command is used to take a short break after completing a Pomodoro session. 
Typically, this break lasts for 5 minutes, allowing you to recharge before starting your next work session. 
This command encourages you to step away from your work, stretch, or relax for a moment, which can help maintain productivity over time.`
)

const (
	HistoryUse   = "history"
	HistoryShort = ""
	HistoryLong  = ""
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

// Records name
const (
	PomodoroRecord = "🍅 pomodoro:"
	BreakRecord    = "☕️ break:"
)

// Sounds name
const (
	PomodoroSound = "Submarine"
	BreakSound    = "Blow"
)
