package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func maxI(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func minI(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func collision(points []Point, rect1_p1 Point, rect1_p2 Point, rect1_i int, rect1_j int) bool {
	for i, p1 := range points {
		j := i + 1
		// wrap around
		if j >= len(points) {
			j = 0
		}

		if i == rect1_i || i == rect1_j || j == rect1_i || j == rect1_j {
			continue
		}

		p2 := points[j]

		r1_top := minI(rect1_p1.y, rect1_p2.y)
		r1_bottom := maxI(rect1_p1.y, rect1_p2.y)
		r1_left := minI(rect1_p1.x, rect1_p2.x)
		r1_right := maxI(rect1_p1.x, rect1_p2.x)

		r2_top := minI(p1.y, p2.y)
		r2_bottom := maxI(p1.y, p2.y)
		r2_left := minI(p1.x, p2.x)
		r2_right := maxI(p1.x, p2.x)

		if r1_top < r2_bottom && r1_bottom > r2_top && r1_left < r2_right && r1_right > r2_left {
			return true
		}

	}
	return false
}

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
	points := make([]Point, 0)

	for _, l := range lines {
		coordsStr := strings.Split(l, ",")
		x, _ := strconv.Atoi(coordsStr[0])
		y, _ := strconv.Atoi(coordsStr[1])

		points = append(points, Point{x, y})
	}

	getArea := func(p1 Point, p2 Point) int {
		h := 1 + int(math.Abs(float64(p1.y)-float64(p2.y)))
		w := 1 + int(math.Abs(float64(p1.x)-float64(p2.x)))
		return h * w
	}

	max := 0
	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}

			a := getArea(p1, p2)
			if max > a {
				continue
			}

			if isPartTwo && collision(points, p1, p2, i, j) {
				continue
			}

			if max < a {
				max = a
			}
		}
	}
	fmt.Println(max)
}
