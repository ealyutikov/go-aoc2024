package day02

// https://adventofcode.com/2024/day/2

const (
	MinIncrement = 1
	MaxIncrement = 3
)

func partOne(levels [][]int) (result int) {
	for _, line := range levels {
		if isSafe(line) {
			result++
		}
	}

	return result
}

func isSafe(line []int) bool {
	isIncrease, isDecrease := true, true
	for i := 0; i < len(line)-1; i++ {
		if line[i]-line[i+1] < MinIncrement || line[i]-line[i+1] > MaxIncrement {
			isIncrease = false
		}

		if line[i+1]-line[i] < MinIncrement || line[i+1]-line[i] > MaxIncrement {
			isDecrease = false
		}
	}

	return isIncrease || isDecrease
}
