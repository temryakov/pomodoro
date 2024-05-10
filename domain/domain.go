package domain

type IStatus interface {
	Start(chan int)
	Sound()
	GetFinishDescription() string

	// SelectOption() int
}
