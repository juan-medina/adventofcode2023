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

package _test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/juan-medina/adventofcode2023/internal/structs"
	"github.com/juan-medina/adventofcode2023/internal/utils"
)

func TestSolver(s structs.Solver, day int, part int) error {

	wd, _ := os.Getwd()
	os.Chdir(path.Join(wd, "../../internal/_test"))
	defer os.Chdir(wd)

	i, err := utils.GetFile(day, part, "input")

	if err != nil {
		return fmt.Errorf("fail to get file: %v", err)
	}

	got, err := s.Solve(i, part)

	if err != nil {
		return fmt.Errorf("fail to get result: %v", err)
	}

	expect, err := utils.GetFile(day, part, "solution")

	if err != nil {
		return fmt.Errorf("fail to get file: %v", err)
	}

	if !utils.CompareStringSlices(got, expect) {
		return fmt.Errorf("solving error, expect: %v, got %v", expect, got)
	}
	return nil
}

func init() {
	fmt.Println("Executing global test setup...")
	cwd, err := os.Getwd()
	if err == nil {
		np := path.Join(cwd, "../../..")
		os.Chdir(np)
	}
}

func TestGlobalSetup(t *testing.T) {
	// This test function will not run actual tests,
	// but the init() function will be executed
	// before any tests when `go test` is called
}
