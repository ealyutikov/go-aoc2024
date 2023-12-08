package day02

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string

func TestSolution(t *testing.T) {
	println(solution(input))
	assert.Equal(t, 2439, solution(input))
}

func TestMain(m *testing.M) {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	m.Run()
}
