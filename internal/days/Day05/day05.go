package Day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juan-medina/adventofcode2023/internal/structs"
)

type Day05 struct {
	structs.BasicSolver
}

func New() Day05 {
	return Day05{}
}

func (obj Day05) Run(day int, part int) error {
	return obj.BasicSolver.BasicRun(obj, day, part)
}

func (obj Day05) Solve(input []string, part int) ([]string, error) {
	result := []string{}

	a := getAlmanac(input)

	a = setToLocation(a)

	m := lowestLocation(a)

	result = append(result, strconv.Itoa(m))

	return result, nil
}

func lowestLocation(a almanac) int {
	l := a.seeds[0].location
	for i := range a.seeds {
		s := a.seeds[i]
		if s.location < l {
			l = s.location
		}
	}

	return l
}

func setToLocation(a almanac) almanac {
	r := a

	for i := range r.seeds {
		s := r.seeds[i]

		fmt.Printf("Seed %v", s.id)
		for j := range r.maps {
			m := r.maps[j]
			fmt.Printf(" %v", m.to)

			for z := range m.ranges {
				r := m.ranges[z]

				t := r.from + r.length
				if s.location >= r.from && s.location <= t {
					s.location = r.to + (s.location - r.from)
					break
				}
			}
			r.seeds[i].location = s.location
			fmt.Printf(" %v", r.seeds[i].location)
		}

		fmt.Println()
	}
	return r
}

func getAlmanac(input []string) almanac {
	r := almanac{}

	r.seeds = getSeeds(input[0])
	r.maps = getMaps(input[2:])

	return r
}

func getMaps(input []string) []maps {
	ms := []maps{}

	m := maps{
		ranges: []ranges{},
		from:   "",
		to:     "",
	}

	for i := range input {
		l := input[i]

		// reset map
		if l == "" {
			ms = append(ms, m)
			m = maps{
				ranges: []ranges{},
				from:   "",
				to:     "",
			}
			continue
		}

		// check if we are waiting for a map start
		if m.from == "" {
			sl := l[0 : len(l)-5]
			t := strings.Split(sl, "-to-")
			m.from = t[0]
			m.to = t[1]
			m.ranges = []ranges{}
		} else {
			r := ranges{}
			t := strings.Split(l, " ")

			r.to, _ = strconv.Atoi(t[0])
			r.from, _ = strconv.Atoi(t[1])
			r.length, _ = strconv.Atoi(t[2])

			m.ranges = append(m.ranges, r)
		}
	}

	// add last map if needed
	if m.from != "" {
		ms = append(ms, m)
	}

	return ms
}

func getSeeds(input string) []seed {
	s := []seed{}

	n := input[7:]

	t := strings.Split(n, " ")

	for i := range t {
		d, _ := strconv.Atoi(t[i])
		ns := seed{
			id:       d,
			location: d,
		}
		s = append(s, ns)
	}

	return s
}

type ranges struct {
	to     int
	from   int
	length int
}

type maps struct {
	ranges []ranges
	from   string
	to     string
}

type seed struct {
	id       int
	location int
}

type almanac struct {
	seeds []seed
	maps  []maps
}
