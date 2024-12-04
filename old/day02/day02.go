package day02

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	rCount = 12
	gCount = 13
	bCount = 14
)

var (
	regRed   = regexp.MustCompile(`\d+\sred`)
	regGreen = regexp.MustCompile(`\d+\sgreen`)
	regBlue  = regexp.MustCompile(`\d+\sblue`)
	rexNum   = regexp.MustCompile(`\d+`)
)

func solutionA(input []string) (sum int) {
	game := parseGame(input)
	for id, sets := range game {
		if !check(sets) {
			continue
		}
		sum = sum + id
	}

	return sum
}

func solutionB(input []string) (sum int) {
	game := parseGame(input)
	for _, sets := range game {
		ps := power(sets)
		res := ps.red * ps.green * ps.blue
		sum = sum + res
	}

	return sum
}

func check(sets []Set) bool {
	for _, s := range sets {
		if s.red > rCount || s.green > gCount || s.blue > bCount {
			return false
		}
	}

	return true
}

func parseGame(input []string) Game {
	game := make(Game, len(input))
	for i, line := range input {
		game[i+1] = parseSets(strings.Split(line, ":")[1])
	}

	return game
}

// 1 blue, 1 red; 10 red; 8 red, 1 blue, 1 green; 1 green, 5 blue
func parseSets(line string) []Set {
	sets := strings.Split(line, ";")
	result := make([]Set, len(sets))

	for _, s := range sets {
		r, g, b := parseRGB(s)
		result = append(result, Set{r, g, b})
	}

	return result
}

// 8 red, 1 blue, 1 green
func parseRGB(line string) (red, green, blue int) {
	red, _ = strconv.Atoi(
		strings.Split(regRed.FindString(line), " ")[0],
	)
	green, _ = strconv.Atoi(
		strings.Split(regGreen.FindString(line), " ")[0],
	)
	blue, _ = strconv.Atoi(
		strings.Split(regBlue.FindString(line), " ")[0],
	)

	return red, green, blue
}
