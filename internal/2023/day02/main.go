package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"hvpaiva.dev/aoc-in-go/pkg/goaoc"
	"hvpaiva.dev/aoc-in-go/pkg/parser"
	"hvpaiva.dev/aoc-in-go/pkg/timed"
)

//go:embed input.txt
var input string

const name = "cube conundrum"

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
	var games []Game
	lines := parser.ToLines(input)

	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	gamesAllowed := filterGamesAllowed(games, []Rule{
		{Color: Red, Count: 12},
		{Color: Green, Count: 13},
		{Color: Blue, Count: 14},
	})

	return sumIds(gamesAllowed)
}

func partTwo(input string) int {
	var games []Game
	var sum int

	lines := parser.ToLines(input)

	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	for _, game := range games {
		minSet := game.MinSet()
		sum += minSet.Power()
	}

	return sum
}

type Game struct {
	Id      int
	Subsets []Subset
}

func (g Game) Allowed(rules []Rule) bool {
	for _, subset := range g.Subsets {
		if !subset.Allowed(rules) {
			return false
		}
	}

	return true
}

type Rule struct {
	Color Color
	Count int
}

func NewGame(line string) Game {
	var subsets []Subset

	gameHeader := strings.Split(line, ": ")

	sets := strings.Split(gameHeader[1], "; ")

	for _, set := range sets {
		subsets = append(subsets, NewSubset(set))
	}

	gameId := strings.Split(gameHeader[0], " ")[1]

	id, err := strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}

	return Game{
		Id:      id,
		Subsets: subsets,
	}
}

func (g Game) MinSet() Subset {
	minSet := make(map[Color]int, 3)

	for _, subset := range g.Subsets {
		for color, count := range subset {
			if minSet[color] == 0 || count > minSet[color] {
				minSet[color] = count
			}
		}
	}

	return minSet
}

type Subset map[Color]int

func NewSubset(set string) Subset {
	ss := make(map[Color]int, 3)

	sets := strings.Split(set, ", ")

	for _, st := range sets {
		unit := strings.Split(st, " ")
		i, err := strconv.Atoi(unit[0])
		if err != nil {
			panic(err)
		}
		ss[Color(unit[1])] = +i
	}

	return ss
}

func (s Subset) Allowed(rules []Rule) bool {
	for _, rule := range rules {
		if s[rule.Color] > rule.Count {
			return false
		}
	}
	return true
}

func (s Subset) Power() int {
	var power = 1

	for _, count := range s {
		power *= count
	}

	return power
}

type Color string

var (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

func filterGamesAllowed(games []Game, rules []Rule) []Game {
	var allowed []Game

	for _, game := range games {
		if game.Allowed(rules) {
			allowed = append(allowed, game)
		}
	}

	return allowed
}

func sumIds(games []Game) int {
	var sum int

	for _, game := range games {
		sum += game.Id
	}

	return sum
}
