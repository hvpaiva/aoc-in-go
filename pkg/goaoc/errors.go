package goaoc

import (
	"errors"
	"fmt"
)

type ErrInvalidPart struct {
	Part int
}

func (e ErrInvalidPart) Error() string {
	return fmt.Sprintf("invalid part: %d. The valid parts are (1/2)", e.Part)
}

var ErrInvalidPartType = errors.New("invalid part type. The part type allowed is int")

var ErrMissingPart = errors.New("no part specified, please provide a valid part")

type ErrIORead struct {
	Err error
}

func (e ErrIORead) Error() string {
	return fmt.Sprintf("failed to read input: %v", e.Err)
}

func (e ErrIORead) Unwrap() error {
	return e.Err
}

type ErrIOWrite struct {
	Err error
}

func (e ErrIOWrite) Error() string {
	return fmt.Sprintf("failed to write input: %v", e.Err)
}

func (e ErrIOWrite) Unwrap() error {
	return e.Err
}

var ErrNoSolution = errors.New("no solution found")

type ErrNoSolutionForPart struct {
	Part Part
}

func (e ErrNoSolutionForPart) Error() string {
	return fmt.Sprintf("no solution found for part %d", e.Part)
}
