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

package Day01

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/juan-medina/adventofcode2023/internal/structs"
)

type Day01 struct {
	structs.BasicSolver
}

func (obj Day01) Run(day int, part int) error {
	return obj.BasicSolver.BasicRun(obj, day, part)
}

var stringDigits = map[string]int {
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

func getNumber(line string, index int) (int, error) {
	lineLength := len(line)
	digit := line[index]
	if digit >= '0' && digit <= '9' {
		return int(digit - '0'), nil
	}

	for key := range stringDigits {		
		keyLength := len(key)
		lastIndex := index+keyLength		

		if (lastIndex) < lineLength {
			token := line[index:lastIndex]
			if token == key {
				return stringDigits[key], nil
			}
		}
	}

	return 0, errors.New("no number found")
}

func findFirstDigit(line string) int {
	for i := range line {
		number, err := getNumber(line, i)
		if err== nil {
			return number
		}
	}
	return 0
}

func findLastDigit(line string) int {
	len := len(line)
	for i := range line {
		number, err := getNumber(line, len-i-1)
		if err== nil {
			return number
		}
	}

	return 0
}

func (obj Day01) Solve(input []string, part int) ([]string, error) {
	result := []string{}

	total := 0
	for i := range input {

		line := input[i]
		first := findFirstDigit(line)
		last := findLastDigit(line)
		totalLine := (first * 10) + last

		fmt.Printf("%3d %s\n", totalLine, line)

		total = total + totalLine
	}

	fmt.Println()

	result = append(result, strconv.Itoa(total))

	return result, nil
}

func New() Day01 {
	return Day01{}
}