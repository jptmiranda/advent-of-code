package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ogInput [][]string
var input [][]string
var resultsA []int
var resultsB []int
var regex = regexp.MustCompile(`[^\w.]`)

func init() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	splitInput := strings.Split(string(data), "\n")

	for _, lineString := range splitInput {
		var line []string

		for _, r := range lineString {
			line = append(line, string(r))
		}

		input = append(input, line)
	}

	ogInput = input
}

func main() {
	for i, line := range input {
		for j, char := range line {
			if regex.MatchString(char) {
				getResult(char, j, i)
			}
		}
	}

	// Part A
	sum := 0
	for _, n := range resultsA {
		sum += n
	}
	fmt.Printf("Part A Results: %d\n", sum)

	// Part B
	sum = 0
	for _, n := range resultsB {
		sum += n
	}
	fmt.Printf("Part B Results: %d\n", sum)
}

func getResult(symbol string, x int, y int) {
	var numbersAround []int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if j == x && i == y {
				continue
			}

			_, err := strconv.Atoi(input[i][j])
			// number found
			if err == nil {
				fullNumber, _ := strconv.Atoi(buildNumber(j, i))
				numbersAround = append(numbersAround, fullNumber)
			}
		}
	}

	sum := 0

	for _, n := range numbersAround {
		sum += n
	}

	resultsA = append(resultsA, sum)

	if symbol == "*" && len(numbersAround) == 2 {
		resultsB = append(resultsB, numbersAround[0]*numbersAround[1])
	}
}

func buildNumber(x int, y int) string {
	var fullNumber []string
	startPosition := x

	for {
		if startPosition-1 < 0 {
			break
		}

		_, err := strconv.Atoi(input[y][startPosition-1])
		if err != nil {
			break
		}

		startPosition--
	}

	for {
		if startPosition > len(input[y])-1 {
			break
		}

		_, err := strconv.Atoi(input[y][startPosition])
		if err != nil {
			break
		}

		fullNumber = append(fullNumber, input[y][startPosition])
		input[y][startPosition] = "."
		startPosition++
	}

	return strings.Join(fullNumber, "")
}
