package domain

type PomodoroBreaker interface {
	StartDescription() string
	FinishDescription() string
	Sound()

	// SelectOption() int
}
