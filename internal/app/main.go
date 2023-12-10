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

	Days "github.com/juan-medina/adventofcode2023/internal/days"
)

func main() {

	day := flag.Int("day", 0, "what day to run")
	part := flag.Int("part", 1, "what part to run 1 or 2")

	flag.Parse()

	if *day == 0 || *part < 1 || *part > 2 {
		flag.Usage()
		return
	}

	days := Days.New()

	if *day > len(days) {
		fmt.Printf("No enough examples implement for running day: %d", *day)
		return
	}

	fmt.Printf("running : day: %v [part: %v]\n", *day, *part)

	err := days[*day-1].Run(*day, *part)
	if err != nil {
		fmt.Printf("error running example : %v", err.Error())
		os.Exit(1)
	}
}
