package main

import (
	"fmt"
	"math"
)

var pruned int
var foundOnce bool
var numPossible int
var foundOnceResult []int

// result[i]=j; means queen at i-th row is placed at j-th column.
// Queens placed at (x1, y1) and (x2,y2)
// x1==x2 means same rows, y1==y2 means same columns, |x2-x1|==|y2-y1| means
// they are placed in diagonals.
func promising(x2 int, y2 int, result []int) bool {
	// This function will check if queen can be placed (x2,y2), or we can
	// say that Can queen at x2 row is placed at y2 column.
	// for finding the column for x2 row, we will check all the columns for
	// all the rows till x2-1.
	for i := 0; i < x2; i++ {
		//result[i] == y2 => columns are same
		//|i - x2| == |result[i] - y2| => in diagonals.
		if (result[i] == y2) || (math.Abs(float64(i)-float64(x2)) == math.Abs(float64(result[i])-float64(y2))) {
			return false
		}
	}
	return true
}

func placeQueens(x int, size int, result []int) {
	for i := 0; i < size; i++ {
		//check if queen at xth row can be placed at i-th column.
		if promising(x, i, result) {
			result[x] = i // place the queen at this position.
			if x == size-1 {
				//foundOnce is a global bool variable sinice the function is rescursive
				if !foundOnce {
					fmt.Println("Order of", size, " queens", result)
					foundOnce = true
					for index, element := range result {
						foundOnceResult[index] = element
					}
				}
				numPossible++
			}
			placeQueens(x+1, size, result)
		} else {
			//If the move is not promising, prune the node and count using a global variable
			pruned++
		}
	}
}

func main() {
	n := []int{4, 8, 10, 12} //Orders for n-queens
	for _, element := range n {
		pruned = 0
		foundOnce = false
		numPossible = 0
		foundOnceResult = make([]int, element)
		result := make([]int, element)
		placeQueens(0, element, result)
		for i := 0; i < element; i++ { //row
			for j := 0; j < element; j++ { //column
				if foundOnceResult[i] != j {
					fmt.Print("| |")
				} else {
					fmt.Print("|Q|")
				}
			}
			fmt.Print("\n")
		}
		fmt.Println("Number of possibilties: ", numPossible)
		fmt.Println("Number of nodes pruned: ", pruned)
	}
}
