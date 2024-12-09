package day04

import "strings"

//https://adventofcode.com/2024/day/4

const (
	SAMX = "SAMX"
	XMAS = "XMAS"
)

func partOne(lines []string) (result int) {
	for i := 0; i < len(lines); i++ {
		result += strings.Count(lines[i], XMAS)
		result += strings.Count(lines[i], SAMX)
	}

	return result
}

func partTwo(_ []string) (result int) {
	return result
}
