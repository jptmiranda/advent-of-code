package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []map[string]int
}

func (g *Game) Valid(red int, green int, blue int) bool {
	valid := true

	for _, set := range g.sets {
		if set["red"] > red || set["green"] > green || set["blue"] > blue {
			valid = false
		}
	}

	return valid
}

func (g *Game) MinimumCubesPower() int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, set := range g.sets {
		if set["red"] > maxRed {
			maxRed = set["red"]
		}
		if set["green"] > maxGreen {
			maxGreen = set["green"]
		}
		if set["blue"] > maxBlue {
			maxBlue = set["blue"]
		}
	}

	return maxRed * maxGreen * maxBlue
}

var input string

func init() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input = string(data)
}

func getGameId(game string) int {
	r, _ := regexp.Compile(`^Game (\d+):`)
	id, _ := strconv.Atoi(r.FindStringSubmatch(game)[1])

	return id
}

func getGameSets(game string) []map[string]int {
	r := regexp.MustCompile(`^Game (\d+):`)

	allSets := strings.Trim(r.Split(game, 2)[1], " ")
	splitSets := strings.Split(allSets, ";")

	var sets []map[string]int
	for _, set := range splitSets {
		splitSet := strings.Split(set, ",")

		set := make(map[string]int)
		for _, s := range splitSet {
			r := regexp.MustCompile(`(\d+)\s*([a-zA-Z]+)`)
			match := r.FindStringSubmatch(s)
			quantity, _ := strconv.Atoi(match[1])
			color := match[2]
			set[color] = quantity
		}

		sets = append(sets, set)
	}

	return sets
}

func main() {
	gamesInput := strings.Split(input, "\n")
	var games []Game

	for _, game := range gamesInput {
		id := getGameId(game)
		sets := getGameSets(game)

		games = append(games, Game{
			id,
			sets,
		})
	}

	validGamesIdSum := 0
	for _, game := range games {
		if game.Valid(12, 13, 14) {
			validGamesIdSum += game.id
		}
	}

	gamePowerSum := 0
	for _, game := range games {
		gamePowerSum += game.MinimumCubesPower()
	}

	fmt.Println("id sum of valid games: ", validGamesIdSum)
	fmt.Println("sum of the power of each game: ", gamePowerSum)
}
