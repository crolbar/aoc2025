package main

import (
	"fmt"
	"os"
	"strings"
)

func dfs(lines []string, nLines int, x int, y int) int {
	res := 0

	for i := y; i < nLines; i++ {
		// used splitter, skip
		if lines[i][x] == '#' {
			break
		}
		if lines[i][x] == '^' {
			// mark splitter as used
			{
				b := []byte(lines[i])
				b[x] = '#'
				lines[i] = string(b)
			}
			res += dfs(lines, nLines, x+1, i)
			res += dfs(lines, nLines, x-1, i)
			res += 1
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

	var (
		lines  = strings.Split(strings.Trim(string(c), "\n"), "\n")
		nLines = len(lines)
		sx     = 0
		sy     = 0
	)
	_ = lines

	(func() {
		for i := 0; i < nLines; i++ {
			for j := 0; j < len(lines[0]); j++ {
				if lines[i][j] == 'S' {
					sy = i
					sx = j
					return
				}
			}
		}
	})()

	fmt.Println("out: ", dfs(lines, nLines, sx, sy))
}
