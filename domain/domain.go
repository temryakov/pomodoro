package domain

type IStatus interface {
	Start(chan int)
	Sound()
	FinishDescription() string

	// SelectOption() int
}
