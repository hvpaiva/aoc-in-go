package timed

import (
	"fmt"
	"time"

	"hvpaiva.dev/aoc-in-go/pkg/goaoc"
)

func Record(name, input string, action goaoc.Solution) int {
	start := time.Now()
	res := action(input)

	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)

	return res
}

func RecordSolution(name string, action goaoc.Solution) goaoc.Solution {
	return func(input string) int {
		return Record(name, input, action)
	}
}
