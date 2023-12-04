package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	playedNumbers  []int
	count          int
}

var cards []Card
var wonCards []Card
var totalCards int = 0

func init() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	splitInput := strings.Split(string(data), "\n")
	regex := regexp.MustCompile(`Card\s+(\d+):\s*([\d\s]+)\|\s*([\d\s]+)`)

	for _, line := range splitInput {
		matches := regex.FindAllStringSubmatch(line, -1)

		id, _ := strconv.Atoi(matches[0][1])

		cards = append(cards, Card{
			id:             id,
			winningNumbers: getNumbers(matches[0][2]),
			playedNumbers:  getNumbers(matches[0][3]),
			count:          1,
		})
	}

	wonCards = cards
	totalCards = len(wonCards)
}

func getNumbers(input string) []int {
	var numbers []int

	splitNumbers := strings.Split(input, " ")

	for _, s := range splitNumbers {
		n, err := strconv.Atoi(strings.Trim(s, " "))
		if err == nil {
			numbers = append(numbers, n)
		}
	}

	return numbers
}

func main() {
	sum := 0
	for _, c := range cards {
		sum += c.value()
	}
	for _, c := range wonCards {
		c.checkWin()
	}
	fmt.Println(sum)
	fmt.Println(totalCards)
}

func (c *Card) value() int {
	value := 0

	for _, w := range c.winningNumbers {
		for _, p := range c.playedNumbers {
			if w == p {
				if value == 0 {
					value = 1
					continue
				}

				value *= 2
			}
		}
	}

	return value
}

func (c *Card) checkWin() {
	if c.count > 0 {
		count := 0

		for _, w := range c.winningNumbers {
			for _, p := range c.playedNumbers {
				if w == p {
					count++
				}
			}
		}

		if count > 0 {
			c.count--
			for i := c.id; i < c.id+count; i++ {
				wonCards[i].count++
				totalCards++
			}
			c.checkWin()
		}
	}
}
