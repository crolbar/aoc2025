package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	idx int
	jol int
}
type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }

// i is child, j is parent. the child is going upwards the heap if res is true
func (h MaxHeap) Less(i, j int) bool {
	if h[i].jol == h[j].jol {
		return h[i].idx < h[j].idx
	}
	return h[i].jol > h[j].jol
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
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

	conts := strings.Trim(string(c), "\n")

	res := 0
	for _, bank := range strings.Split(conts, "\n") {
		if isPartTwo {
			res += part2(bank)
			continue
		}
		res += part1(bank)
	}

	fmt.Println(res)
}

// finds the biggest number to the right of idx (with a bigger index than idx)
// d is the depth of the recursion
// return of -1 means didn't find anything, 0 means we reached the end (d == 13)
func dfs(copyH MaxHeap, idx int, d int) int {
	if d == 13 {
		return 0
	}

	h := copyH
	removed := make([]Pair, 0)

	for h.Len() > 0 {
		p := heap.Pop(&h).(Pair)
		// fmt.Printf("[d%d] %d, %d\n", d, p.idx, p.jol)

		// skip numbers to the left
		if p.idx < idx {
			continue
		}

		heapToPass := append(MaxHeap(nil), (h)...)
		{
			if len(removed) > 0 {
				for _, v := range removed {
					heap.Push(&heapToPass, v)
				}
			}
		}

		dfsRes := dfs(heapToPass, p.idx, d+1)
		if dfsRes == -1 {
			removed = append(removed, p)
			continue
		}

		curr := p.jol * int(math.Pow(10, float64(12-d)))
		return curr + dfsRes
	}

	return -1
}

func part2(bank string) int {
	h := &MaxHeap{}
	heap.Init(h)

	for i, b := range bank {
		heap.Push(h, Pair{jol: int(b - '0'), idx: i})
	}

	res := dfs(*h, 0, 1)
	return res
}

/*






















 */

func part1(bank string) int {
	h := &MaxHeap{}
	heap.Init(h)
	// fmt.Println(bank)

	for i, b := range bank {
		heap.Push(h, Pair{jol: int(b - '0'), idx: i})
	}

	for h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		// fmt.Println(p.idx, p.jol)

		if p.idx+1 >= len(bank) {
			continue
		}

		mr := 0
		for _, b := range bank[p.idx+1:] {
			n := int(b - '0')
			if n > mr {
				mr = n
			}
		}

		bankJolt := (p.jol * 10) + mr
		// fmt.Println(bankJolt)
		return bankJolt
	}
	panic("did not found jolt")
}

/*
 813783562874281
 | |||||||||| |

 837835628748
 878356287428















	987654321111
	900000000000 - 1
	 80000000000 - 2
			....
			   1 - 12

*/
