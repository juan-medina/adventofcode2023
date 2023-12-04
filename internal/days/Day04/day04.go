package Day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juan-medina/adventofcode2023/internal/structs"
)

type Day04 struct {
	structs.BasicSolver
}

func New() Day04 {
	return Day04{}
}

func (obj Day04) Run(day int, part int, test bool) error {
	return obj.BasicSolver.BasicRun(obj, day, part, test)
}

func (obj Day04) Solve(input []string, part int) ([]string, error) {
	result := []string{}

	g := getGame(input)

	t := getScore(g)

	result = append(result, strconv.Itoa(t))

	return result, nil
}

func getScore(g game) int {
	total := 0

	for i := range g {
		total = total + getCardScore(g[i])
	}

	return total
}

func getCardScore(s scratchCard) int {
	total := 0

	fmt.Printf("Card %v winning numbers: ", s.id)
	for i := range s.play {
		n := s.play[i]
		for j := range s.winning {
			w := s.winning[j]
			if n == w {
				if total == 0 {
					total = 1
				} else {
					total = total + total
				}
				fmt.Printf("%v ", j)
				break
			}
		}
	}
	if total == 0 {
		fmt.Printf("none ")
	}
	fmt.Printf("= %v points\n", total)

	return total
}

func getGame(input []string) game {
	g := game{}

	for i := range input {
		c := getScratchCard(input[i])
		g = append(g, c)
	}

	return g
}

func getScratchCard(line string) scratchCard {
	c := scratchCard{}

	t := strings.Split(line, ": ")

	sid := strings.TrimSpace(t[0][5:])
	c.id, _ = strconv.Atoi(sid)

	t = strings.Split(t[1], " | ")

	c.winning = getNumbers(t[0])
	c.play = getNumbers(t[1])

	return c
}

func getNumbers(part string) numbers {
	r := numbers{}

	l := len(part)

	for i := 0; i < l; i += 3 {
		sn := strings.TrimSpace(part[i : i+2])
		n, _ := strconv.Atoi(sn)
		r = append(r, n)
	}

	return r
}

type numbers []int

type scratchCard struct {
	id      int
	winning numbers
	play    numbers
}

type game []scratchCard
