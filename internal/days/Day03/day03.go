package Day03

import (
	"fmt"
	"strconv"

	"github.com/juan-medina/adventofcode2023/internal/structs"
)

type Day03 struct {
	structs.BasicSolver
}

func New() Day03 {
	return Day03{}
}

func (obj Day03) Run(day int, part int, test bool) error {
	return obj.BasicSolver.BasicRun(obj, day, part, test)
}

func (obj Day03) Solve(input []string, part int) ([]string, error) {
	result := []string{}

	e := getEngine(input)

	total := 0

	if part == 1 {
		total = sumOfParts(e)
	} else {
		total = gearRatio(e)
	}

	result = append(result, strconv.Itoa(total))

	return result, nil
}

type element struct {
	row int
	col int
}

type possibleNumber struct {
	star  element
	end   element
	value int
}

type symbol struct {
	element
	char byte
}

type engine struct {
	elements []possibleNumber
	symbols  []symbol
}

func getEngine(lines []string) engine {
	e := engine{
		symbols:  getSymbols(lines),
		elements: getPossibleNumbers(lines),
	}

	return e
}

func getSymbols(lines []string) []symbol {
	s := []symbol{}

	for i := range lines {
		for j := range lines[i] {
			c := lines[i][j]
			if isSymbol(c) {
				e := symbol{element: element{
					row: i,
					col: j,
				},
					char: c,
				}
				s = append(s, e)
			}
		}
	}

	return s
}

func isSymbol(b byte) bool {
	return !isNum(b) && !isPeriod(b)
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isPeriod(b byte) bool {
	return b == '.'
}

func sumOfParts(e engine) int {

	total := 0

	for i := range e.elements {
		v := false
		p := e.elements[i]
		for j := range e.symbols {
			s := e.symbols[j]
			v = adjacent(p, s.element)
			if v {
				fmt.Printf("element %v (%v,%v) collide with symbol '%v'at (%v,%v)\n", p.value, p.star.row, p.star.col, string(s.char), s.row, s.col)
				break
			}
		}
		if v {
			fmt.Printf("element value is %v\n", p.value)
			total = total + p.value
		} else {
			fmt.Printf("element %v (%v,%v) does not collide with any symbol\n", p.value, p.star.row, p.star.col)
		}

	}

	return total
}

func gearRatio(e engine) int {

	total := 0
	for i := range e.symbols {

		s := e.symbols[i]

		if s.char == '*' {
			totalParts := 0
			ratio := 1
			for j := range e.elements {
				p := e.elements[j]
				if adjacent(p, s.element) {
					totalParts = totalParts + 1
					ratio = ratio * p.value
					if totalParts == 2 {
						fmt.Printf("gear at (%v,%v) has a ratio of: %v\n", s.row, s.col, ratio)
						total += ratio
						break
					}
				}
			}
		}
	}
	return total
}

func adjacent(p possibleNumber, s element) bool {

	// if they are in the same row just the sides
	if s.row == p.star.row {
		if s.col == p.star.col-1 || s.col == p.end.col+1 {
			return true
		}
	}

	// up or bottom row, should be on the line +-1
	if s.row == p.star.row-1 || s.row == p.star.row+1 {
		if s.col >= p.star.col-1 && s.col <= p.end.col+1 {
			return true
		}
	}

	return false
}

func getPossibleNumbers(lines []string) []possibleNumber {
	p := []possibleNumber{}

	c := possibleNumber{
		star: element{
			row: -1,
			col: -1,
		},
		end: element{
			row: -1,
			col: -1,
		},
		value: 0,
	}

	for i := range lines {
		// new row reset
		c.star.row = -1
		c.star.col = -1
		c.end.row = -1
		c.end.col = -1
		c.value = 0

		for j := range lines[i] {

			d := lines[i][j]

			// we search for an start
			if c.star.row == -1 {
				if isNum(d) {
					c.star.row = i
					c.star.col = j
					c.value = int(d - '0')
				}
				// we search for an end
			} else {
				//if is a num
				if isNum(d) {
					// continue adding
					c.value = (c.value * 10) + int(d-'0')
					// if we reach the end
					if j == len(lines[i])-1 {
						c.end.row = i
						c.end.col = j

						p = append(p, c)

						c.star.row = -1
						c.star.col = -1
						c.end.row = -1
						c.end.col = -1
						c.value = 0
					}
					// we end this number
				} else {
					c.end.row = i
					c.end.col = j - 1

					p = append(p, c)

					c.star.row = -1
					c.star.col = -1
					c.end.row = -1
					c.end.col = -1
					c.value = 0
				}
			}
		}
	}

	return p
}
