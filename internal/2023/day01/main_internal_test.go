package main

import "testing"

func TestPartOne(t *testing.T) {
	input = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	want := 142

	got := partOne(input)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}

}

func TestPartTwo(t *testing.T) {
	input = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\noneight"
	want := 299

	got := partTwo(input)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
