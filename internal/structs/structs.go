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

package structs

import (
	"github.com/juan-medina/adventofcode2023/internal/utils"
)

type DaySolver interface {
	Run(day int, part int) error
}

type Solver interface {
	Solve(input []string, part int) ([]string, error)
}

type BasicSolver struct {
	DaySolver
}

func (obj BasicSolver) BasicRun(solver Solver, day int, part int) error {
	inputData, err := utils.GetFile(day, part, "input")

	if err != nil {
		return err
	}

	utils.OutputStringSlice("input data:", inputData)

	solved, err := solver.Solve(inputData, part)
	if err == nil {
		utils.OutputStringSlice("solution:", solved)
	}

	return err
}
