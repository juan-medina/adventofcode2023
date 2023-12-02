/*
 * Copyright (c) 2023 Juan Antonio Medina Iglesias
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 */

package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juan-medina/adventofcode2023/internal/structs"
)

type Day02 struct {
	structs.BasicSolver
}

func (obj Day02) Run(day int, part int, test bool) error {
	return obj.BasicSolver.BasicRun(obj, day, part, test)
}

type ballDraw struct {
	balls int
	color string
}

type bagDraw []ballDraw

type ballGame struct {
	id      int
	bagDraw []bagDraw
}

func getGameId(idPart string) int {
	id, _ := strconv.Atoi(strings.Split(idPart, " ")[1])
	return id
}

func getDrawsInBag(bagPart string) bagDraw {
	result := bagDraw{}

	draws := strings.Split(bagPart, ",")

	for i := range draws {
		parts := strings.Split(draws[i], " ")

		draw := ballDraw{}
		draw.balls, _ = strconv.Atoi(parts[1])
		draw.color = parts[2]

		result = append(result, draw)
	}

	return result
}

func getBallDraws(ballPart string) []bagDraw {
	result := []bagDraw{}

	draws := strings.Split(ballPart, ";")

	for i := range draws {
		bags := []ballDraw{}
		bagDraw := getDrawsInBag(draws[i])
		bags = append(bags, bagDraw...)
		result = append(result, bags)
	}

	return result
}

func getBallGame(line string) ballGame {
	parts := strings.Split(line, ":")

	game := ballGame{}

	game.id = getGameId(parts[0])
	game.bagDraw = getBallDraws(parts[1])

	return game
}

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

var maxBallsPerGame = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func countGame(game ballGame) int {
	for i := range game.bagDraw {
		draws := game.bagDraw[i]

		for j := range draws {
			draw := draws[j]
			max := maxBallsPerGame[draw.color]
			if draw.balls > max {
				fmt.Printf("game id: %v is not valid, because it has too many %v balls: %v, max: %v\n", game.id, draw.color, draw.balls, max)
				return 0
			}
		}
	}

	fmt.Printf("game id: %v is valid\n", game.id)

	return game.id
}

func (obj Day02) Solve(input []string, part int) ([]string, error) {
	games := []ballGame{}
	result := []string{}

	for i := range input {
		line := input[i]
		game := getBallGame(line)
		games = append(games, game)
	}

	total := 0
	for i := range games {
		total = total + countGame(games[i])
	}

	result = append(result, strconv.Itoa(total))

	return result, nil
}
