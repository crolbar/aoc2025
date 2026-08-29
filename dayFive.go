package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	c, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	isPartTwo := false
	if len(os.Args) > 2 {
		isPartTwo = os.Args[2] == "2"
	}
	_ = isPartTwo

	lines := strings.Split(strings.Trim(string(c), "\n"), "\n")

	ranges := [][]int{}

	i := 0

	for len(lines[i]) != 0 {
		nums := strings.Split(lines[i], "-")
		n, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, []int{n, n2})
		i++
	}

	i++

	out := 0
	for i < len(lines) {
		for _, fresh_range := range ranges {
			n, _ := strconv.Atoi(lines[i])

			if n >= fresh_range[0] && n <= fresh_range[1] {
				out += 1
				break
			}
		}
		i++
	}

	fmt.Println("ans", out)

	// PART TWO
	// a bit of spaggeti but does the job
	ranges_no_overlaps := [][]int{}
	var has_overlap = func(r1 []int, r2 []int) bool {
		if r1[0] >= r2[0] && r1[0] <= r2[1] {
			return true
		}
		if r1[1] >= r2[0] && r1[1] <= r2[1] {
			return true
		}

		// r2 is in r1
		if r2[0] >= r1[0] && r2[0] <= r1[1] {
			return true
		}
		if r2[1] >= r1[0] && r2[1] <= r1[1] {
			return true
		}
		return false
	}

	var remove = func(_ranges [][]int, r []int) [][]int {
		for i, _r := range _ranges {
			if _r[0] == r[0] && _r[1] == r[1] {
				return append(slices.Clone(_ranges[:i]), _ranges[i+1:]...)
			}
		}
		return _ranges
	}

	var combine = func(_ranges [][]int) [][]int {
		new_ranges := slices.Clone(_ranges)
	start:
		for i, r1 := range new_ranges {
			for j, r2 := range new_ranges {
				if i == j {
					continue
				}

				if has_overlap(r1, r2) {
					new_ranges = append(new_ranges, []int{min(r1[0], r2[0]), max(r1[1], r2[1])})
					nr1 := remove(new_ranges, r1)
					f := remove(nr1, r2)
					new_ranges = f
					goto start
				}
			}
		}
		return new_ranges
	}

	ranges_no_overlaps = combine(ranges)

	out = 0
	for _, fresh := range ranges_no_overlaps {
		out += (fresh[1] - fresh[0]) + 1
	}

	fmt.Println("ans2", out)
}
