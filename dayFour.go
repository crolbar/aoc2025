package main

import (
	"fmt"
	"os"
	"strings"
)

var directions = [][]int{
	{0, 1},   // right
	{1, 1},   // down right
	{1, 0},   // down
	{1, -1},  // down left
	{0, -1},  // left
	{-1, -1}, // up left
	{-1, 0},  // up
	{-1, 1},  // up right
}

func run(lines []string, partTwo bool) int {
	res := 0

	height := len(lines)
	width := len(lines[0])

	// result of one passthrough
	resRun := -1

	for resRun != 0 {
		resRun = 0

		linesCopy := make([]string, len(lines))
		copy(linesCopy, lines)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if linesCopy[i][j] != '@' {
					continue
				}

				paper := 0
				for _, dir := range directions {
					dy := i + dir[0]
					dx := j + dir[1]

					if dx < 0 || dx >= width {
						continue
					}
					if dy < 0 || dy >= height {
						continue
					}

					if linesCopy[dy][dx] == '@' {
						paper += 1
						if paper >= 4 {
							break
						}
					}
				}

				if paper >= 4 {
					continue
				} else {
					lineRunes := []rune(lines[i])
					lineRunes[j] = 'x'
					lines[i] = string(lineRunes)

					resRun += 1
				}
			}
		}

		res += resRun
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println()
		if !partTwo {
			break
		}
	}

	return res
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

	res := run(lines, isPartTwo)

	fmt.Println(res)
}
