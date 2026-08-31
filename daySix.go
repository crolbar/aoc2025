package main

import (
	"fmt"
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

	var (
		lines    = strings.Split(strings.Trim(string(c), "\n"), "\n")
		nLines   = len(lines)
		lastLine = lines[nLines-1]
		lineLen  = len(lastLine)
		out      = 0
		numOps   = 0
	)
	for i := 0; i <= len(lastLine)-1; i++ {
		if lastLine[i] == '+' || lastLine[i] == '*' {
			numOps += 1
		}
	}

	getNextNum := func(str string) int {
		s := strings.TrimSpace(str)

		nextWhiteSpace := len(s)
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				nextWhiteSpace = i
				break
			}
		}

		n, err := strconv.Atoi(s[:nextWhiteSpace])
		if err != nil {
			panic(err)
		}

		return n
	}

	ci := 0
	for range numOps {
		op := lastLine[ci]
		opRes := 0
		if op == '*' {
			opRes = 1
		}

		// fmt.Println(string(op))
		for ri := nLines - 2; ri >= 0; ri-- {
			n := getNextNum(lines[ri][ci:])
			if op == '+' {
				opRes += n
			} else if op == '*' {
				opRes *= n
			}
		}

		nextOpIdx := lineLen - 1
		for i := ci + 1; i < lineLen; i++ {
			if lastLine[i] == '+' || lastLine[i] == '*' {
				nextOpIdx = i
				break
			}
		}

		ci = nextOpIdx
		out += opRes

		// fmt.Println("res", opRes)
		// fmt.Println()
	}

	fmt.Printf("out: %d\n", out)
}
