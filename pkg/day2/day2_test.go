package day2

import "testing"

var testCases = []struct {
	input string
	exp   bool
	expB  bool
}{
	{
		input: "1-3 a: abcde",
		exp:   true,
		expB:  true,
	},
	{
		input: "1-3 b: cdefg",
		exp:   false,
		expB:  false,
	},
	{
		input: "2-9 c: ccccccccc",
		exp:   true,
		expB:  false,
	},
}

func TestIsValid(t *testing.T) {
	for _, tc := range testCases {
		if res, _ := IsValid(tc.input); res != tc.exp {
			t.Errorf("IsValid(%s): want %t, got %t", tc.input, tc.exp, res)
		}
	}
}

func TestIsValidB(t *testing.T) {
	for _, tc := range testCases {
		if res, _ := IsValidB(tc.input); res != tc.expB {
			t.Errorf("IsValid(%s): want %t, got %t", tc.input, tc.expB, res)
		}
	}
}

func TestDay2(t *testing.T) {
	t.Logf("\nSolution: %v\nSolutionB: %v", Solve(input), SolveB(input))
}
