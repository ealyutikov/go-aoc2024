package day01

import "slices"

func solutionA(listA, listB []int) (result int) {
	slices.Sort(listA)
	slices.Sort(listB)

	for i := 0; i < len(listA); i++ {
		result += diff(listA[i], listB[i])
	}

	return result
}

func solutionB(listA, listB []int) (result int) {
	for a := 0; a < len(listA); a++ {
		count := 0
		for b := 0; b < len(listB); b++ {
			if listA[a] == listB[b] {
				count++
			}
		}

		result += listA[a] * count
	}

	return result
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
