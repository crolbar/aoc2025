package main

import (
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

	n := 50
	password := 0
	rotations := strings.Split(string(c), "\n")

	for _, rotation := range rotations {
		if len(rotation) == 0 {
			continue
		}
		direction := rotation[0]

		raw := 0
		numRotations, err := strconv.ParseInt(rotation[1:], 10, 32)
		if err != nil {
			panic(err)
		}

		switch direction {
		case 'L':
			raw = n - int(numRotations)

			tmp := raw
			for tmp < 0 {
				tmp = 100 + tmp
				if isPartTwo {
					if n == 0 && tmp >= 0 {
						continue
					}
					password += 1
				}
			}

			if tmp < 0 || tmp > 99 {
				panic("rotation out of range")
			}

			n = tmp
		case 'R':
			raw = n + int(numRotations)

			tmp := raw
			for tmp > 99 {
				tmp = tmp - 100
				if isPartTwo {
					if tmp == 0 {
						continue
					}
					password += 1
				}
			}

			if tmp < 0 || tmp > 99 {
				panic("rotation out of range")
			}

			n = tmp
		default:
			panic("invalid direction")
		}

		println("rotated at: ", n)
		if n == 0 {
			password += 1
		}
	}

	// println(string(c))
	println(password)
}

/*

	50 - 55 = -5

	100 + -5 = 95

	---

	50 - 155 = -105

	100 + -105 = -5
	100 + -5 = 95


	--


	50 + 55 = 105

	100 - 105 = 5

	--

	50 + 155 = 205

	205 - 100 = 105
	105 - 100 = 5

*/
