package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string
var numbersInFull = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func init() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input = string(data)
}

func main() {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		value := getCalibrationValue(line)
		sum += value
	}

	fmt.Println(sum)
}

func getIfNumber(r rune) (int, error) {
	return strconv.Atoi(string(r))
}

func getSlice(line string, start int, count int) string {
	slice := ""

	for i := start; i < start+count; i++ {
		if i <= len(line)-1 {
			slice += string(line[i])
		}
	}

	return slice
}

func getCalibrationValue(line string) int {
	var numbers []int

	for i := 0; i < len(line); i++ {
		if n, err := getIfNumber(rune(line[i])); err == nil {
			numbers = append(numbers, n)
			continue
		}

		for k, v := range numbersInFull {
			slice := getSlice(line, i, len(k))

			if slice == k {
				numbers = append(numbers, v)
			}
		}
	}

	joinedNumbers := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
	value, _ := strconv.Atoi(joinedNumbers)

	fmt.Printf("%d|", value)

	return value
}
