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
	"sort"
	"testing"
)

func TestNewRanges(t *testing.T) {

	ranges := NewRanges(2, 5)

	if ranges.from != 2 {
		t.Errorf("Expected 1 from, got %v", ranges.from)
	}

	if ranges.length != 5 {
		t.Errorf("Expected 5 length, got %v", ranges.length)
	}
}

func CompareIntSlices(slice []int, expect []int) bool {
	if len(slice) != len(expect) {
		return false
	}

	// sort arrays before comparing
	sort.Ints(slice)
	sort.Ints(expect)

	for i := range slice {
		if slice[i] != expect[i] {
			return false
		}
	}
	return true
}

func TestRangesToSlice(t *testing.T) {

	ranges := NewRanges(2, 6)

	slice := ranges.ToSlice()

	if len(slice) != 6 {
		t.Errorf("Expected 5 length, got %v", len(slice))
	}

	expect := []int{2, 3, 4, 5, 6, 7}

	if !CompareIntSlices(slice, expect) {
		t.Errorf("Expected %v, got %v", expect, slice)
	}

	ranges = NewRanges(8, 0)
	expect = []int{}

	if !CompareIntSlices(ranges.ToSlice(), expect) {
		t.Errorf("Expected %v, got %v", expect, ranges.ToSlice())
	}
}

func TestNewRangesFromPairsSlice(t *testing.T) {

	slice := []int{1, 2, 3, 4}

	ranges := NewRangesFromPairsSlice(slice)

	if len(ranges) != 2 {
		t.Errorf("Expected 2 length, got %v", len(ranges))
	}

	slice1 := ranges[0].ToSlice()
	expect1 := []int{1, 2}

	if !CompareIntSlices(slice1, expect1) {
		t.Errorf("Expected %v, got %v", expect1, slice1)
	}

	slice2 := ranges[1].ToSlice()
	expect2 := []int{3, 4, 5, 6}

	if !CompareIntSlices(slice2, expect2) {
		t.Errorf("Expected %v, got %v", expect2, slice2)
	}
}

func TestRangeInOut(t *testing.T) {
	testCases := []struct {
		name   string
		range1 Ranges
		range2 Ranges
		in     []Ranges
		out    []Ranges
	}{
		{
			name:   "FULL OUT LEFT",
			range1: Ranges{from: 1, length: 2},
			range2: Ranges{from: 3, length: 5},
			in:     []Ranges{},
			out:    []Ranges{{1, 2}},
		},
		{
			name:   "FULL OUT RIGHT",
			range1: Ranges{from: 8, length: 2},
			range2: Ranges{from: 3, length: 5},
			in:     []Ranges{},
			out:    []Ranges{{8, 2}},
		},
		{
			name:   "FULL OTHER COVERED BY THIS",
			range1: Ranges{from: 1, length: 7},
			range2: Ranges{from: 3, length: 3},
			in:     []Ranges{{3, 3}},
			out:    []Ranges{{1, 2}, {6, 2}},
		},
		{
			name:   "FULL OTHER ARE SAME",
			range1: Ranges{from: 3, length: 3},
			range2: Ranges{from: 3, length: 3},
			in:     []Ranges{{3, 3}},
			out:    []Ranges{},
		},
		{
			name:   "FULL THIS COVER BY OTHER",
			range1: Ranges{from: 3, length: 3},
			range2: Ranges{from: 1, length: 7},
			in:     []Ranges{{3, 3}},
			out:    []Ranges{},
		},
		{
			name:   "PARTIAL THIS LEFT",
			range1: Ranges{from: 1, length: 4},
			range2: Ranges{from: 3, length: 3},
			in:     []Ranges{{3, 2}},
			out:    []Ranges{{1, 2}},
		},
		{
			name:   "PARTIAL THIS RIGHT",
			range1: Ranges{from: 6, length: 4},
			range2: Ranges{from: 1, length: 7},
			in:     []Ranges{{6, 2}},
			out:    []Ranges{{8, 2}},
		},
		{
			name:   "THIS SINGLE ITEM IN",
			range1: Ranges{from: 1, length: 1},
			range2: Ranges{from: 1, length: 1},
			in:     []Ranges{{1, 1}},
			out:    []Ranges{},
		},
		{
			name:   "THIS SINGLE ITEM OUT LEFT",
			range1: Ranges{from: 1, length: 1},
			range2: Ranges{from: 3, length: 1},
			in:     []Ranges{},
			out:    []Ranges{{1, 1}},
		},
		{
			name:   "THIS SINGLE ITEM OUT RIGHT",
			range1: Ranges{from: 5, length: 1},
			range2: Ranges{from: 3, length: 1},
			in:     []Ranges{},
			out:    []Ranges{{5, 1}},
		},
		{
			name:   "OTHER SINGLE ITEM OUT LEFT",
			range1: Ranges{from: 2, length: 2},
			range2: Ranges{from: 1, length: 1},
			in:     []Ranges{},
			out:    []Ranges{{2, 2}},
		},
		{
			name:   "OTHER SINGLE ITEM OUT RIGHT",
			range1: Ranges{from: 2, length: 2},
			range2: Ranges{from: 3, length: 1},
			in:     []Ranges{{3, 1}},
			out:    []Ranges{{2, 1}},
		},
		{
			name:   "OTHER SINGLE ITEM IN",
			range1: Ranges{from: 1, length: 3},
			range2: Ranges{from: 2, length: 1},
			in:     []Ranges{{2, 1}},
			out:    []Ranges{{1, 1}, {3, 1}},
		},
        {
            name : "FAR THIS LEFT",
            range1: Ranges{from: 1, length: 3},
            range2: Ranges{from: 4, length: 3},
            in:    []Ranges{},
            out:   []Ranges{{1, 3}},
        },
        {
            name : "FAR THIS LEFT SINGLE",
            range1: Ranges{from: 1, length: 1},
            range2: Ranges{from: 4, length: 3},
            in:    []Ranges{},
            out:   []Ranges{{1, 1}},
        },
        {
            name : "FAR THIS LEFT SINGLE OTHER",
            range1: Ranges{from: 1, length: 3},
            range2: Ranges{from: 6, length: 1},
            in:    []Ranges{},
            out:   []Ranges{{1, 3}},
        },        
        {
            name : "FAR THIS RIGHT",
            range1: Ranges{from: 7, length: 3},
            range2: Ranges{from: 4, length: 3},
            in:    []Ranges{},
            out:   []Ranges{{7, 3}},
        },
        {
            name : "FAR THIS LEFT SINGLE",
            range1: Ranges{from: 7, length: 1},
            range2: Ranges{from: 4, length: 3},
            in:    []Ranges{},
            out:   []Ranges{{7, 1}},
        },
        {
            name : "FAR THIS LEFT SINGLE OTHER",
            range1: Ranges{from: 7, length: 3},
            range2: Ranges{from: 4, length: 1},
            in:    []Ranges{},
            out:   []Ranges{{7, 3}},
        },        
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d: %s In Check", i+1, testCase.name), func(t *testing.T) {
			got := testCase.range1.In(testCase.range2)
			if len(got) != len(testCase.in) {
				t.Fatalf("Case %d: Expected %v in, got %v, range 1 = %v, range 2 = %v",
					i+1, testCase.in, got, testCase.range1.ToSlice(), testCase.range2.ToSlice())
			} else {
				for j := range got {
					if !CompareIntSlices(got[j].ToSlice(), testCase.in[j].ToSlice()) {
						t.Fatalf("Case %d: Expected %v in, got %v, range 1 = %v, range 2 = %v",
							i+1, testCase.in, got, testCase.range1.ToSlice(), testCase.range2.ToSlice())
					}
				}
			}
		})
		t.Run(fmt.Sprintf("Case %d: %s Out Check", i+1, testCase.name), func(t *testing.T) {
			got := testCase.range1.Out(testCase.range2)
			if len(got) != len(testCase.out) {
				t.Fatalf("Case %d: Expected %v out, got %v, range 1 = %v, range 2 = %v",
					i+1, testCase.out, got, testCase.range1.ToSlice(), testCase.range2.ToSlice())
			} else {
				for j := range got {
					if !CompareIntSlices(got[j].ToSlice(), testCase.out[j].ToSlice()) {
						t.Fatalf("Case %d: Expected %v out, got %v, range 1 = %v, range 2 = %v",
							i+1, testCase.out, got, testCase.range1.ToSlice(), testCase.range2.ToSlice())
					}
				}
			}
		})

	}
}
