package main

import (
	"fmt"
	"math"
	"os"
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
	points := make([][]int, 0)

	for _, l := range lines {
		coordsStr := strings.Split(l, ",")
		x, _ := strconv.Atoi(coordsStr[0])
		y, _ := strconv.Atoi(coordsStr[1])
		points = append(points, []int{x, y})
	}

	getArea := func(p1 []int, p2 []int) int {
		h := 1 + int(math.Abs(float64(p1[1])-float64(p2[1])))
		w := 1 + int(math.Abs(float64(p1[0])-float64(p2[0])))
		return h * w
	}

	max := 0
	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}
			a := getArea(p1, p2)
			if max < a {
				max = a
			}
		}
	}
	fmt.Println(max)
}
