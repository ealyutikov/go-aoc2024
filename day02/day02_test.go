package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var levels [][]int

func TestMain(m *testing.M) {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		ints := make([]int, len(line))
		for i, s := range line {
			ints[i], _ = strconv.Atoi(s)
		}

		levels = append(levels, ints)
	}

	m.Run()
}

func TestPartOne(t *testing.T) {
	fmt.Printf("%v", levels)
}

func TestPartTwo(t *testing.T) {
	fmt.Printf("%v", levels)
}
