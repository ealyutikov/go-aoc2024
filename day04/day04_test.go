package day04

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputA []string
var inputB []string

func TestSolutionA(t *testing.T) {
	assert.Equal(t, 13768818, solutionB(inputB))
}

func TestMain(m *testing.M) {
	file, err := os.Open("input_b.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		inputB = append(inputB, scanner.Text())
	}

	m.Run()
}
