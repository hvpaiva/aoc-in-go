package main

import "testing"

func TestPartOne(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"Game 1": {
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  1,
		},
		"Game 2": {
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  2,
		},
		"Game 3": {
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:  0,
		},
		"All Games": {
			input: `
					Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
					Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
					Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
					Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
					Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`,
			want: 8,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := partOne(tc.input)

			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}

}

func TestPartTwo(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  int
	}{
		"Game 1": {
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  48,
		},
		"Game 2": {
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  12,
		},
		"Game 3": {
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:  1560,
		},
		"All Games": {
			input: `
					Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
					Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
					Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
					Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
					Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`,
			want: 2286,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got := partTwo(tc.input)

			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}

}
