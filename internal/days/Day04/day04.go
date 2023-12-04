package Day04

import (
	"fmt"
	"math"
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

	t := 0
	if part == 1 {
		t = getScore(g)
	} else {
		t = getTotalCards(g)
	}

	result = append(result, strconv.Itoa(t))

	return result, nil
}

func getScore(g game) int {
	total := 0

	for i := range g {
		wins := getTotalWin(g[i])
		if wins >= 2 {
			wins = int(math.Pow(2.0, float64(wins-1)))
		}
		fmt.Printf(" = %v points\n", wins)
		total = total + wins
	}

	return total
}

func getTotalCards(g game) int {
	total := 0
	m := len(g)

	for i := range g {
		wins := getTotalWin(g[i])
		fmt.Println()

		f := i + 1
		limit := min(f+wins, m)
		for j := f; j < limit; j++ {
			a := g[i].total
			fmt.Printf("adding %v cards to %v\n", a, j+1)
			g[j].total += a
		}
	}

	for i := range g {
		fmt.Printf("card %v has %v total\n", g[i].id, g[i].total)
		total += g[i].total
	}

	return total
}

func getTotalWin(s scratchCard) int {
	total := 0

	fmt.Printf("Card %v winning numbers: ", s.id)
	for i := range s.play {
		n := s.play[i]
		for j := range s.winning {
			w := s.winning[j]
			if n == w {
				total = total + 1
				fmt.Printf("%v ", j)
				break
			}
		}
	}
	if total == 0 {
		fmt.Printf("none ")
	}
	fmt.Printf("= %v wins", total)

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
	c.total = 1

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
	total   int
}

type game []scratchCard
