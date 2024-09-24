package goaoc

type Solution func(string) int

type Challenge struct {
	Input   string
	PartOne Solution
	PartTwo Solution
	Execute Part
}

type Part int

const (
	PartOne Part = iota + 1
	PartTwo
)

func (c Challenge) RunE(options ...RunOption) error {
	var opts runOptions
	if err := injectOptions(&opts, options...); err != nil {
		return err
	}

	part, err := c.whichPartToRun(opts)
	if err != nil {
		return err
	}

	result, err := c.execute(part)
	if err != nil {
		return err
	}

	if err = opts.manager.OutputResults(result); err != nil {
		return err
	}

	return nil
}

func (c Challenge) Run(options ...RunOption) {
	err := c.RunE(options...)
	if err != nil {
		panic(err)
	}
}

func (c Challenge) whichPartToRun(opts runOptions) (Part, error) {
	if c.Execute == 1 || c.Execute == 2 {
		return c.Execute, nil
	}

	if c.PartOne != nil && c.PartTwo == nil {
		return PartOne, nil
	}

	if c.PartOne == nil && c.PartTwo != nil {
		return PartTwo, nil
	}

	var partInput int
	if err := opts.manager.Input(&partInput); err != nil {
		return 0, err
	}

	return Part(partInput), nil
}

func (c Challenge) execute(part Part) (result int, err error) {
	if c.PartOne == nil && c.PartTwo == nil {
		return 0, ErrNoSolution
	}

	solutions := map[Part]Solution{
		PartOne: c.PartOne,
		PartTwo: c.PartTwo,
	}

	solution, ok := solutions[part]
	if !ok {
		return 0, ErrNoSolutionForPart{Part: part}
	}

	return solution(c.Input), nil
}
