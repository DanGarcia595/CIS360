package main

import (
	"fmt"
)

func main() {
	m := 10
	myArray := []int{7, 8, 1, 4, 5, 0, 3, -3, 10, 7, 8, 9, 5, 2, 4, 6, 8, 10, 2345, 5, 12, 2, 1, 0, -100000, 2}
	fmt.Println("Input Array: ", myArray)
	fmt.Println("Size of list (m): ", m)
	fmt.Println(mList(m, myArray))
}

func mList(m int, narray []int) []int {
	if m >= len(narray) {
		return narray
	}
	var finalArray []int
	finalArray = make([]int, m)
	for index, _ := range finalArray {
		finalArray[index] = narray[index]
	}

	var temp int
	var compareVar int
	var i int
	for i = m; i < len(narray); i++ {
		compareVar = narray[i]
		for index, element := range finalArray {
			if compareVar < element {
				temp = compareVar
				compareVar = finalArray[index]
				finalArray[index] = temp
			}
		}
	}
	return finalArray
}
