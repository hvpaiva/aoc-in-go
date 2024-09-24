package goaoc

type RunOption func(options *runOptions) error

type runOptions struct {
	manager IOManager
}

type IOManager interface {
	Input(part *int) error
	OutputResults(result int) error
}

func WithManager(manager IOManager) RunOption {
	return func(options *runOptions) error {
		options.manager = manager

		return nil
	}
}

func injectOptions(opts *runOptions, options ...RunOption) error {
	for _, option := range options {
		_ = option(opts)
	}

	if opts.manager == nil {
		opts.manager = DefaultConsoleManager{Env: defaultConsoleEnv}
	}

	return nil
}
