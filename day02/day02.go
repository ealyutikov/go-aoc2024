package day02

// https://adventofcode.com/2024/day/2

const (
	MinIncrement = 1
	MaxIncrement = 3
)

func partOne(levels [][]int) (result int) {
	for _, line := range levels {
		if isSafeOne(line) {
			result++
		}
	}

	return result
}

func partTwo(levels [][]int) (result int) {
	for _, line := range levels {
		if isSafeTwo(line) {
			result++
		}
	}

	return result
}

func isSafeOne(line []int) bool {
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

func isSafeTwo(line []int) bool {
	for i := 0; i < len(line); i++ {
		if isSafeOne(remove(line, i)) {
			return true
		}
	}

	return false
}

func remove(line []int, idx int) (res []int) {
	res = append(res, line[:idx]...)
	return append(res, line[idx+1:]...)
}
