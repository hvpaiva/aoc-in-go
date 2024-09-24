package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"hvpaiva.dev/aoc-in-go/pkg/goaoc"
	"hvpaiva.dev/aoc-in-go/pkg/parser"
	"hvpaiva.dev/aoc-in-go/pkg/timed"
)

//go:embed input.txt
var input string

const name = "trebuchet"

var partName = func(part goaoc.Part) string {
	return fmt.Sprintf("%s part %d", name, part)
}

var challenge = goaoc.Challenge{
	Input:   input,
	PartOne: timed.RecordSolution(partName(goaoc.PartOne), partOne),
	PartTwo: timed.RecordSolution(partName(goaoc.PartTwo), partTwo),
	Execute: goaoc.PartTwo,
}

func main() {
	challenge.Run()
}

func partOne(input string) int {
	lines := parser.ToLines(input)

	var sum int

	for _, line := range lines {
		firstDigit, lastDigit := -1, -1
		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == -1 {
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0')
			}
		}

		if firstDigit != -1 && lastDigit != -1 {
			sum += firstDigit*10 + lastDigit
		}
	}

	return sum
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	var sum int

	wordToDigit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range lines {
		var firstDigit, lastDigit int
		foundFirst := false
		var wordBuffer strings.Builder

		i := 0
		for i < len(line) {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				num := int(char - '0')
				if !foundFirst {
					firstDigit = num
					foundFirst = true
				}
				lastDigit = num
				wordBuffer.Reset()
				i++
			} else if unicode.IsLetter(char) {
				wordBuffer.WriteRune(char)
				word := wordBuffer.String()

				for key, val := range wordToDigit {
					if strings.HasSuffix(word, key) {
						if !foundFirst {
							firstDigit = val
							foundFirst = true
						}
						lastDigit = val

						processedPart := len(key)
						i = i - processedPart + 1
						wordBuffer.Reset()
					}
				}
				i++
			} else {
				wordBuffer.Reset()
				i++
			}
		}

		if firstDigit != -1 && lastDigit != -1 {
			sum += firstDigit*10 + lastDigit
		}
	}

	return sum
}
