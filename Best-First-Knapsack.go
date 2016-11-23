package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value rune
	Right *Tree
}

type Queue []*node

func (q *Queue) Push(n *node) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n *node) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

type node struct {
	level  int
	profit int
	weight int
	bound  float64
}

func main() {
	n := 5
	p := []int{20, 30, 35, 12, 3}
	w := []int{2, 5, 7, 3, 1}
	W := 13
	maxprofit, nodes := knapsack3(n, p, w, W)
	fmt.Println("for n: ", n, "\np: ", p, "\nw: ", w, "\nW: ", W)
	fmt.Println("The max profit is ", maxprofit, "\nThe number of nodes is ", nodes)
	fmt.Println("---------------------------------------------------------------")
	n = 6
	p = []int{40, 48, 108, 8, 15, 16}
	w = []int{4, 6, 18, 2, 5, 8}
	W = 21
	fmt.Println("for n: ", n, "\np: ", p, "\nw: ", w, "\nW: ", W)
	maxprofit, nodes = knapsack3(n, p, w, W)
	fmt.Println("The max profit is ", maxprofit, "\nThe number of nodes is ", nodes)
	fmt.Println("---------------------------------------------------------------")
	n = 3
	p = []int{20, 25, 2}
	w = []int{3, 1, 2}
	W = 5
	fmt.Println("for n: ", n, "\np: ", p, "\nw: ", w, "\nW: ", W)
	maxprofit, nodes = knapsack3(n, p, w, W)
	fmt.Println("The max profit is ", maxprofit, "\nThe number of nodes is ", nodes)

}

func knapsack3(n int, p []int, w []int, W int) (float64, int) {
	PQ := &Queue{}
	v := &node{-1, 0, 0, 0}
	numnodes := 0
	maxprofit := 0.0
	v.bound = bound(v, W, n, w, p)
	PQ.Push(v)
	for PQ.Len() > 0 {
		v := PQ.Pop()
		if v.bound > maxprofit {
			u := &node{0, 0, 0, 0}
			u.level = v.level + 1
			u.weight = v.weight + w[u.level]
			u.profit = v.profit + p[u.level]

			if (u.weight <= W) && (float64(u.profit) > maxprofit) {
				maxprofit = float64(u.profit)
			}

			u.bound = bound(u, W, n, w, p)
			if u.bound > maxprofit {
				node1 := u
				PQ.Push(node1)
				numnodes++
			}
			u2 := &node{0, 0, 0, 0}
			u2.weight = v.weight
			u2.level = u.level
			u2.profit = v.profit
			u2.bound = bound(u2, W, n, w, p)

			if u2.bound > maxprofit {
				node2 := u2
				PQ.Push(node2)
				numnodes++
			}
		}
	}
	return maxprofit, numnodes
}

func bound(u *node, W int, n int, w []int, p []int) float64 {
	j := -1
	k := 0
	totalweight := 0

	var result float64
	if u.weight > W {
		return 0
	} else {
		result = float64(u.profit)
		j = u.level + 1
		totalweight = u.weight
		for (j < n) && (totalweight+w[j] <= W) {
			totalweight += w[j]
			result += float64(p[j])
			j++
		}
		k = j
		if k < n {
			result += (float64(W) - float64(totalweight)) * (float64(p[k]) / float64(w[k]))
		}
		return result
	}

}
