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

package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func GetFile(day int, part int, name string) ([]string, error) {
	filename := fmt.Sprintf("data/day_%02d_part_%02d_%v.txt", day, part, name)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func CompareStringSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func OutputStringSlice(label string, strings []string) {

	fmt.Println(label)
	fmt.Println()
	for i := range strings {
		fmt.Println(strings[i])
	}
	fmt.Println()
}

type Ranges struct {
	from   int
	length int
}

func NewRanges(from int, length int) Ranges {
	return Ranges{
		from:   from,
		length: length,
	}
}

func NewRangesFromPairsSlice(slice []int) []Ranges {
	//the slice is of the form [from, to, from, to, ...]
	ranges := make([]Ranges, 0)

	for i := 0; i < len(slice); i += 2 {
		ranges = append(ranges, NewRanges(slice[i], slice[i+1]))
	}

	return ranges
}

func (r Ranges) ToSlice() []int {
	if r.length <= 0 {
		return []int{}
	}
	slice := make([]int, r.length)
	for i := range slice {
		slice[i] = r.from + i
	}
	return slice
}

//go:inline
func (r Ranges) End() int {
	return r.from + r.length
}

func (r Ranges) In(other Ranges) []Ranges {
	start := max(r.from, other.from)
	end := min(r.End(), other.End())

	if start >= end {
		return nil
	}

	return []Ranges{{start, end - start}}
}

func (r Ranges) Out(other Ranges) []Ranges {
	out := []Ranges{}

	// check if this range is not full inside other
	if r.from < other.from || r.End() > other.End() {
		// check if this range is not colliding with other
		if r.End() <= other.from || r.from >= other.End() {
			// add this range to the output
			out = append(out, r)
		} else {

			// check if other is inside this
			if other.from >= r.from && other.End() < r.End() {
				// we are in the middle, we need to split in two
				left := Ranges{r.from, other.from - r.from}
				right := Ranges{other.End(), r.End() - (other.End())}

				out = append(out, left, right)
			} else {
				// check if we collide with the left side
				if r.from < other.from {
					out = append(out, Ranges{r.from, other.from - r.from})
					return out
				}

				// check if we collide with the right side
				if r.End() > other.from {
					out = append(out, Ranges{other.End(), r.End() - (other.End())})
					return out
				}
			}
		}
	}

	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mergeRanges(input []Ranges) []Ranges {
	// if we have less than 2 ranges, we can't merge them, return the input
	if len(input) <= 1 {
		return input
	}

	// Sort the input slice based on the 'from' value of each range
	sort.Slice(input, func(i, j int) bool {
		return input[i].from < input[j].from
	})

	merged := make([]Ranges, 0)

	// Add the first range to the merged slice
	merged = append(merged, input[0])

	// iterate over the rest of the ranges
	for i := 1; i < len(input); i++ {
		// get the last range in the merged slice
		lastMerged := &merged[len(merged)-1]

		//check if the current range is full inside the last merged range
		if input[i].from >= lastMerged.from && input[i].End() <= lastMerged.End() {
			continue
		}

		// check if the current range collides with the last merged range or is adjacent
		if lastMerged.End() >= input[i].from {
			// calculate the new end
			end := max(input[i].End(), lastMerged.End())
			// update the last merged range with the new end
			lastMerged.length = end - lastMerged.from
		} else {
			// add the current range to the merged slice
			merged = append(merged, input[i])
		}
	}

	return merged
}
