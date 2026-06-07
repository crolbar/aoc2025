package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func print_bytes(n int) {
	for n > 0 {
		fmt.Print(n & 1) // prints bits LSB first
		n >>= 1
	}
	println()
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

	conts := strings.Trim(string(c), "\n")

	rangesStr := strings.Split(conts, ",")

	res := 0

	for _, v := range rangesStr {
		r := strings.Split(v, "-")
		start, _ := strconv.ParseInt(r[0], 10, 64)
		end, _ := strconv.ParseInt(r[1], 10, 64)

		if start > end {
			panic("invalid range")
		}

		for i := int(start); i <= int(end); i++ {
			numdigits := int(math.Floor(math.Log10(float64(i)))) + 1
			s := strconv.Itoa(i)
			if numdigits%2 == 0 {
				if s[:numdigits/2] == s[numdigits/2:] {
					res += i
					continue
				}
			}

			if !isPartTwo {
				continue
			}

			windowSize := 1
			for windowSize <= numdigits/2 {
				pattern := s[:windowSize]
				invalid := true

				for j := windowSize; j < len(s); j += windowSize {
					end := j + windowSize
					if end <= len(s) && strings.Compare(pattern, s[j:end]) == 0 {
						continue
					}
					invalid = false
				}

				if invalid {
					res += i
					break
				}
				windowSize += 1
			}
		}
	}

	fmt.Println(res)
}
