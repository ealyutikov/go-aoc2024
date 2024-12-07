package day03

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var contents string

func TestMain(m *testing.M) {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents = buf.String()

	m.Run()
}

func TestPartOne(t *testing.T) {
	assert.Equal(t, 472, partOne(contents))
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 0, partTwo(contents))
}
