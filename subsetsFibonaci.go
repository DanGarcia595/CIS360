package main

import (
	"fmt"
	//	"math"
	"time"
)

func main() {

	/*	myArray := [7]int{10, 30, 40, 50, 60, 65, 70}
		for _, element := range myArray {
			start := time.Now()
			fib1(element)
			elapsed := time.Since(start)
			fmt.Println(element, " took: ", elapsed)
			fmt.Println("T/2^(n/2): ", float64(elapsed)/math.Pow(2, float64(element/2)))
		}*/
	myArray2 := [7]int{50, 100, 1000, 10000, 50000, 10000000}
	for _, element := range myArray2 {
		start := time.Now()
		fib2(element)
		elapsed := time.Since(start)
		fmt.Println(element, " took: ", elapsed)
		fmt.Println("T/n: ", float64(elapsed)/float64(element))
	}

}

func Subsets(array []int) {
	for i := 0; i < len(array); i++ {
		if i < len(array)-1 {
			for j := i + 1; j < len(array); j++ {
				if j < len(array)-1 {
					for k := j + 1; k < len(array); k++ {
						fmt.Println(array[i], array[j], array[k])
					}
				}
			}
		}
	}

}

func SmallestLargest(array []int) {
	smallest := array[0]
	largest := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < smallest {
			smallest = array[i]
		} else if array[i] > largest {
			largest = array[i]
		}
	}
	fmt.Println(largest, " ", smallest)

}

func fib1(num int) int {
	if num <= 1 {
		return num
	} else {
		return fib1(num-1) + fib1(num-2)
	}

}

func fib2(n int) int {
	f := make([]int, n+1)

	f[0] = 0
	if n > 0 {
		f[1] = 1
		for i := 2; i <= n; i++ {
			f[i] = f[i-1] + f[i-2]
		}
	}
	return f[n]
}
