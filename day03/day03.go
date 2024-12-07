package day03

import (
	"regexp"
	"strconv"
	"strings"
)

func partOne(contents string) (result int) {
	pattern := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := pattern.FindAllString(contents, -1)

	for i := 0; i < len(matches); i++ {
		result += multiply(matches[i])
	}

	return result
}

func partTwo(contents string) (result int) {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do|don't)\(\)`)
	matches := pattern.FindAllString(contents, -1)

	doMul := true
	for i := 0; i < len(matches); i++ {
		if matches[i] == "do()" {
			doMul = true
			continue
		}

		if matches[i] == "don't()" {
			doMul = false
			continue
		}

		if doMul {
			result += multiply(matches[i])
		}
	}

	return result
}

func multiply(src string) int {
	tmp := strings.ReplaceAll(src, `mul(`, "")
	tmp = strings.ReplaceAll(tmp, `)`, "")
	args := strings.Split(tmp, ",")

	a1, _ := strconv.Atoi(args[0])
	a2, _ := strconv.Atoi(args[1])

	return a1 * a2
}
