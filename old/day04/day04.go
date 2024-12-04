package day04

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var re = regexp.MustCompile(`\d+`)

func solutionA(input []string) (sum int) {
	for _, line := range input {
		res := lo.Intersect(parse(line))
		sum += int(math.Pow(2, float64(len(res)-1)))
	}

	return sum
}

func solutionB(lines []string) int {
	copies := make([]int, len(lines))
	total := 0
	for i, line := range lines {
		copies[i]++
		winners, ticket := parse(line)
		wins := 0
		for _, number := range ticket {
			if lo.Contains(winners, number) {
				wins++
				copies[i+wins] += copies[i]
			}
		}
		total += copies[i]
	}
	
	return total
}

func parse(line string) ([]int, []int) {
	res := strings.Split(strings.Split(line, ":")[1], "|")

	return lo.Map(re.FindAllString(res[0], -1), func(str string, _ int) int {
			res, _ := strconv.Atoi(str)
			return res
		}),
		lo.Map(re.FindAllString(res[1], -1), func(str string, _ int) int {
			res, _ := strconv.Atoi(str)
			return res
		})
}
