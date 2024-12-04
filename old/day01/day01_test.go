package day01

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string

func TestSolution(t *testing.T) {
	assert.Equal(t, 54953, solution(input))
}

func TestSolutionF(t *testing.T) {
	assert.Equal(t, 54953, solutionF(input))
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solution(input)
	}
}

func BenchmarkSolutionF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solutionF(input)
	}
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
