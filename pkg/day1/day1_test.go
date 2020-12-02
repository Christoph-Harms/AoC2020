package day1

import (
	"strings"
	"testing"
)

var testCases = []struct {
	input string
	exp   int
}{
	{
		input: "2019\n1\n142\n4737128\n4824",
		exp:   2019,
	},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		if solution := Solve(tc.input); solution != tc.exp {
			t.Errorf("Solve(\"%s\"): expected %d, got %d", strings.ReplaceAll(tc.input, "\n", " "), tc.exp, solution)
		}
	}
}

var testCasesB = []struct {
	input string
	exp   int
}{
	{
		input: "2017\n1\n2\n142\n4737128\n4824",
		exp:   4034,
	},
}

func TestSolveB(t *testing.T) {
	for _, tc := range testCasesB {
		if solution := SolveB(tc.input); solution != tc.exp {
			t.Errorf("Solve(\"%s\"): expected %d, got %d", strings.ReplaceAll(tc.input, "\n", " "), tc.exp, solution)
		}
	}
}

func TestDay1(t *testing.T) {
	t.Logf("\nSolution: %v\nSolution2: %v", Solve(input), SolveB(input))
}
