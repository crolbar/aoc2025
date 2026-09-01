package main

import (
	"fmt"
	"os"
	"strings"
)

func dfs(lines []string, nLines int, dp [][]int, p1 bool, x int, y int) int {
	for i := y; i < nLines; i++ {
		// stop on used splitters on p1
		if p1 && dp[i][x] == 1 {
			break
		}

		// use cached res of splitter on p2
		if !p1 && dp[i][x] != 0 {
			return dp[i][x]
		}

		if lines[i][x] == '^' {
			// mark splitter as used for p1
			if p1 {
				dp[i][x] = 1
			}

			fmt.Println("s", i+1, x+1)

			res := dfs(lines, nLines, dp, p1, x+1, i) + dfs(lines, nLines, dp, p1, x-1, i)

			if p1 {
				res += 1
			} else
			// cache splitter timelines
			{
				dp[i][x] = res
			}
			return res
		}
	}

	if p1 {
		return 0
	} else {
		return 1
	}
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
		dp     = make([][]int, 0)
		sx     = 0
		sy     = 0
	)
	for i := 0; i < nLines; i++ {
		dp = append(dp, make([]int, len(lines[0])))
	}
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

	fmt.Println("out: ", dfs(lines, nLines, dp, !isPartTwo, sx, sy))
}
