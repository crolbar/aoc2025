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
	points := make([][]float64, 0)

	// circuits[id][0] -> pointer circuit size
	// if circuits[id][1] -> connecte to other circuit
	circuits := make([][]*int, 0)
	// point idx -> circuit id
	circuitsMap := make(map[int]int)

	for _, l := range lines {
		coordsStr := strings.Split(l, ",")
		x, _ := strconv.Atoi(coordsStr[0])
		y, _ := strconv.Atoi(coordsStr[1])
		z, _ := strconv.Atoi(coordsStr[2])
		points = append(points, []float64{float64(x), float64(y), float64(z)})
	}

	for i, p1 := range points {
		// closest point to p1
		closestPIdx := -1
		closestDist := -1.0
		for j, p2 := range points {
			if j == i {
				continue
			}
			xd := math.Abs(p1[0] - p2[0])
			yd := math.Abs(p1[1] - p2[1])
			zd := math.Abs(p1[2] - p2[2])
			dist := math.Sqrt(math.Pow(xd, 2) + math.Pow(yd, 2) + math.Pow(zd, 2))
			if closestDist == -1.0 || dist < closestDist {
				closestPIdx = j
				closestDist = dist
			}
		}
		if closestPIdx == -1 {
			panic("did not found closest")
		}

		var (
			p1CircuitId = -1
			p2CircuitId = -1
			ok          = false
		)
		if p1CircuitId, ok = circuitsMap[i]; !ok {
			p1CircuitId = -1
		}
		if p2CircuitId, ok = circuitsMap[closestPIdx]; !ok {
			p2CircuitId = -1
		}

		fmt.Println("closest to", i, "is", closestPIdx)
		// new circuit
		if p1CircuitId == -1 && p2CircuitId == -1 {
			n := 2
			n2 := 0
			circuits = append(circuits, []*int{&n, &n2})
			id := len(circuits) - 1
			circuitsMap[i] = id
			circuitsMap[closestPIdx] = id
			continue
		}

		// connect p1 to p2 circuit
		if p1CircuitId == -1 {
			circuitsMap[i] = p2CircuitId
			*circuits[p2CircuitId][0] += 1
			continue
		}
		// connect p2 to p1 circuit
		if p2CircuitId == -1 {
			circuitsMap[closestPIdx] = p1CircuitId
			*circuits[p1CircuitId][0] += 1
			continue
		}

		if p2CircuitId == p1CircuitId {
			continue
		}

		// connect p2 circuit to p1 circuit
		*circuits[p1CircuitId][0] += *circuits[p2CircuitId][0]
		circuits[p2CircuitId][0] = circuits[p1CircuitId][0]
		*circuits[p2CircuitId][1] = 1

		// fmt.Println(p1CircuitId, p2CircuitId)
	}

	for i, c := range circuits {
		if *c[1] == 0 {
			fmt.Println(i, ":", *c[0])
		}
	}
}
