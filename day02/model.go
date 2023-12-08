package day02

type Set struct {
	red   int
	green int
	blue  int
}

type Game = map[int][]Set

func power(sets []Set) Set {
	var rMax, gMax, bMax int

	for _, s := range sets {
		rMax = max(s.red, rMax)
		gMax = max(s.green, gMax)
		bMax = max(s.blue, bMax)
	}

	return Set{rMax, gMax, bMax}
}
