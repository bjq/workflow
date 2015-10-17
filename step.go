package workflow

type StepFunc func(context Context) error

type Step struct {
	Label     string
	Run       StepFunc
	IsLast    bool
	DependsOn []*Step
}
