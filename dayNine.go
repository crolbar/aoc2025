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
	safePoints := make(map[Point]bool)

	maxX := math.MaxInt
	minX := math.MaxInt
	maxY := math.MaxInt
	minY := math.MaxInt
	for _, l := range lines {
		coordsStr := strings.Split(l, ",")
		x, _ := strconv.Atoi(coordsStr[0])
		y, _ := strconv.Atoi(coordsStr[1])

		maxX = int(math.Max(float64(maxX), float64(x)))
		minX = int(math.Min(float64(minX), float64(x)))
		maxY = int(math.Max(float64(maxY), float64(y)))
		minY = int(math.Min(float64(minY), float64(y)))

		points = append(points, Point{x, y})
	}
	topLeftP := Point{minX, minY}
	bottomRightP := Point{maxX, maxY}

	fmt.Println(topLeftP, bottomRightP)

	var prevP Point = points[len(points)-1]
	for _, p := range points {

		// outline from prev point to current
		{
			maxX := int(math.Max(float64(p.x), float64(prevP.x)))
			minX := int(math.Min(float64(p.x), float64(prevP.x)))
			maxY := int(math.Max(float64(p.y), float64(prevP.y)))
			minY := int(math.Min(float64(p.y), float64(prevP.y)))
			for x := minX; x <= maxX; x++ {
				safePoints[Point{x, p.y}] = true
			}
			for y := minY; y <= maxY; y++ {
				safePoints[Point{p.x, y}] = true
			}
		}

		safePoints[p] = true
		prevP = p
	}
	fmt.Println("outline")

	// fill in the outline in safePoints
	for y := topLeftP.y; y < bottomRightP.y; y++ {
		flag := false
		isInside := false
		for x := topLeftP.x; x < bottomRightP.x; x++ {
			// make sure we mark isInside when we reach a point like `#.`
			// not just a `#`. this messes up when we have multiple `#`s
			if !flag && safePoints[Point{x, y}] {
				flag = true
			} else if flag && !safePoints[Point{x, y}] {
				isInside = true
				flag = false
			} else if isInside && safePoints[Point{x, y}] {
				isInside = false
			}

			if isInside {
				safePoints[Point{x, y}] = true
			}
		}
	}

	fmt.Println("after fill")

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
			if !safePoints[Point{p1.x, p2.y}] {
				continue
			}
			if !safePoints[Point{p2.x, p1.y}] {
				continue
			}
			a := getArea(p1, p2)
			if max < a {
				max = a
			}
		}
	}
	fmt.Println(max)

	// /*

	//  */

	// t := make([]string, 0)
	// for i := 0; i < 10; i++ {
	// 	var sb strings.Builder
	// 	for j := 0; j < 20; j++ {
	// 		if safePoints[Point{j, i}] {
	// 			sb.WriteByte('#')
	// 		} else {
	// 			sb.WriteByte('.')
	// 		}
	// 	}
	// 	t = append(t, sb.String())
	// }
	// fmt.Println(strings.Join(t, "\n"))
}
