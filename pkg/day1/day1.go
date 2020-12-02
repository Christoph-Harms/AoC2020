package day1

import (
	"strconv"
	"strings"
)

func Solve(input string) int {
	intIn := string2IntSlice(input)

	return solve(intIn)
}

func SolveB(input string) int {
	intIn := string2IntSlice(input)

	return solveB(intIn)
}

func solve(in []int) int {
	isLEQ2020 := func(x int) bool { return x <= 2020 }

	in = filter(in, isLEQ2020)

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i]+in[j] == 2020 {
				return in[i] * in[j]
			}
		}
	}

	return 0
}

func solveB(in []int) int {
	isLEQ2020 := func(x int) bool { return x <= 2020 }

	in = filter(in, isLEQ2020)

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			for k := j + 1; k < len(in); k++ {
				if in[i]+in[j]+in[k] == 2020 {
					return in[i] * in[j] * in[k]
				}
			}
		}
	}

	return 0
}

// filter filters in place without extra memory allocation.
// Stolen from https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
func filter(in []int, keepCondition func(int) bool) []int {
	b := in[:0]
	for _, x := range in {
		if keepCondition(x) {
			b = append(b, x)
		}
	}

	return b
}

func string2IntSlice(in string) []int {
	var out []int

	for _, s := range strings.Split(in, "\n") {
		i, _ := strconv.Atoi(s)
		out = append(out, i)
	}

	return out
}
