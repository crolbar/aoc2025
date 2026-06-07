package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	idx int
	jol int
}
type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].jol > h[j].jol }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
		h := &MaxHeap{}
		heap.Init(h)
		// fmt.Println(bank)

		for i, b := range bank {
			heap.Push(h, Pair{jol: int(b - '0'), idx: i})
		}

		for h.Len() > 0 {
			p := heap.Pop(h).(Pair)

			if p.idx+1 >= len(bank) {
				continue
			}

			mr := 0
			for _, b := range bank[p.idx+1:] {
				n := int(b-'0')
				if n > mr {
					mr = n
				}
			}

			bankJolt := (p.jol*10)+mr
			// fmt.Println(bankJolt)
			res += bankJolt
			break
		}
	}

	fmt.Println(res)
}
