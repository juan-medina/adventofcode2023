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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/juan-medina/adventofcode2023/internal/days/Day01"
	"github.com/juan-medina/adventofcode2023/internal/days/Day02"
	"github.com/juan-medina/adventofcode2023/internal/structs"
)

func main() {

	day := flag.Int("day", 0, "what day to run")
	part := flag.Int("part", 1, "what part to run 1 or 2")
	test := flag.Bool("test", false, "test or normal run")

	flag.Parse()

	if *day == 0 || *part < 1 || *part > 2 {
		flag.Usage()
		return
	}

	var example structs.DaySolver = nil

	examples := []structs.DaySolver{
		Day01.New(),
		Day02.New(),
	}

	if *day > len(examples) {
		fmt.Printf("No enough examples implement for running day: %d", *day)
		return
	}

	example = examples[*day-1]

	if *test {
		fmt.Printf("running : day: %v [part: %v] TEST\n", *day, *part)
	} else {
		fmt.Printf("running : day: %v [part: %v]\n", *day, *part)
	}

	err := example.Run(*day, *part, *test)
	if err != nil {
		fmt.Printf("error running example : %v", err.Error())
		os.Exit(1)
	}
}
