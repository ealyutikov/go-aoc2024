package day04

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var content []string

func TestMain(m *testing.M) {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	m.Run()
}

func TestPartOne(t *testing.T) {
	assert.Equal(t, 472, partOne(content))
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 0, partTwo(content))
}
