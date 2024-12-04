package day01

import (
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

var re = regexp.MustCompile("[0-9]")

func solutionF(input []string) int {
	return lo.Reduce(input, func(acc int, line string, _ int) int {
		nums := re.FindAllString(line, -1)
		num, _ := strconv.Atoi(nums[0] + nums[len(nums)-1])
		return acc + num
	}, 0)
}

func solution(input []string) int {
	acc := 0
	for _, line := range input {
		nums := re.FindAllString(line, -1)
		num, _ := strconv.Atoi(nums[0] + nums[len(nums)-1])
		acc += num
	}

	return acc
}
