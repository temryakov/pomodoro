package domain

type Model struct {
	Use               string
	ShortDescription  string
	LongDescription   string
	StartDescription  string
	FinishDescription string
	Duration          int
	HasFinish         bool
	HasPause          bool
}

type IStatus interface {
	Start(set int)
	Finish(spent int)
	// SelectOption() int
}
