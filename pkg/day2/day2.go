package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) int {
	inSlice := strings.Split(input, "\n")

	numValid := 0
	for _, in := range inSlice {
		if valid, _ := IsValid(in); valid {
			numValid++
		}
	}

	return numValid
}

func SolveB(input string) int {
	inSlice := strings.Split(input, "\n")

	numValid := 0
	for _, in := range inSlice {
		if valid, _ := IsValidB(in); valid {
			numValid++
		}
	}

	return numValid
}

func IsValid(input string) (bool, error) {
	min, max, subject, password, err := parseInput(input)
	if err != nil {
		return false, err
	}

	subjectOccurrences := 0
	for _, r := range password {
		if r == subject {
			subjectOccurrences++
		}
	}

	if subjectOccurrences < min || subjectOccurrences > max {
		return false, nil
	}

	return true, nil
}

func IsValidB(input string) (bool, error) {
	posA, posB, subject, password, err := parseInput(input)
	if err != nil {
		return false, err
	}

	if (rune(password[posA]) == subject) != (rune(password[posB]) == subject) { // XOR
		return true, nil
	}

	return false, nil
}

func parseInput(input string) (int, int, rune, string, error) {
	outerParts := strings.Split(input, ":")

	rule := outerParts[0]
	password := outerParts[1]

	parts := strings.Split(rule, " ")

	minmax := strings.Split(parts[0], "-")

	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return 0, 0, 0, "", err
	}

	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return 0, 0, 0, "", err
	}

	if len(parts[1]) > 1 {
		return 0, 0, 0, "", fmt.Errorf("rule can only refer to a single character, instead got %s", parts[1])
	}

	subject := []rune(parts[1])[0]

	return min, max, subject, password, nil
}
