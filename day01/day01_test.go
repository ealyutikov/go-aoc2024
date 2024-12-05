package day01

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputL, inputR []int

func TestMain(m *testing.M) {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	i := 0
	for scanner.Scan() {
		res, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if (i % 2) == 0 {
			inputL = append(inputL, res)
		} else {
			inputR = append(inputR, res)
		}

		i++
	}

	if len(inputL) != len(inputR) {
		log.Fatal(errors.New("inputL != inputR"))
	}

	m.Run()
}

func TestPartOne(t *testing.T) {
	assert.Equal(t, 2769675, partOne(inputL, inputR))
}

func TestPartTwoLoop(t *testing.T) {
	assert.Equal(t, 24643097, partTwoLoop(inputL, inputR))
}

func TestPartTwoMap(t *testing.T) {
	assert.Equal(t, 24643097, partTwoMap(inputL, inputR))
}
