package day01

import (
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func solution(input []string) int {
	re := regexp.MustCompile("[0-9]")

	return lo.Reduce(input, func(acc int, line string, _ int) int {
		ints := re.FindAllString(line, -1)
		if len(ints) == 0 {
			return acc
		}

		number, _ := strconv.Atoi(ints[0] + ints[len(ints)-1])
		return acc + number
	}, 0)
}
