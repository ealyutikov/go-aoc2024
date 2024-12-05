package day01

import "slices"

func partOne(listA, listB []int) (result int) {
	slices.Sort(listA)
	slices.Sort(listB)

	for i := 0; i < len(listA); i++ {
		result += diff(listA[i], listB[i])
	}

	return result
}

func partTwoLoop(listA, listB []int) (result int) {
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

func partTwoMap(listA, listB []int) (result int) {
	setB := map[int]int{}
	for _, v := range listB {
		setB[v] = setB[v] + 1
	}

	for _, v := range listA {
		result += v * setB[v]
	}

	return result
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
