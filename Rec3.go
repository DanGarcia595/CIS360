package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//	myarray := []int{0, 8, 5, 4, 5, 6}
	//	fmt.Println(DivConqMax(myarray))
	/*num := 17
	fmt.Println("Num ", num)
	fmt.Println(Div3Conq(num, myarray))*/
	NewTowerOfHanoi(10)
	SimulateGame()

}
func DivConqMax(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	length := len(array)
	fmt.Println(array)
	left := DivConqMax(array[0:(length / 2)])
	right := DivConqMax(array[length/2 : length])
	if left > right {
		return left
	} else {
		return right
	}
}

func Div3Conq(num int, array []int) string {
	length := len(array)
	fmt.Println(array)
	var mystring string
	if length > 2 {
		if num > array[(length/3)-1] {
			if num > array[2*(length/3)-1] {
				mystring = Div3Conq(num, array[2*(length/3):length])
			} else {
				mystring = Div3Conq(num, array[(length/3):2*(length/3)])
			}
		} else {
			mystring = Div3Conq(num, array[0:(length/3)])
		}
	} else if length > 1 {
		if array[0] == num || array[1] == num {
			return "Value Found"
		}
	}
	if mystring == "Value Found" {
		return "Value Found"
	} else {
		return "Value Not Found"
	}
}

var SourcePeg *[]int
var TargetPeg *[]int
var HelperPeg *[]int

func NewTowerOfHanoi(numOfDisks int) {
	SourcePeg = getNsortedNum(numOfDisks)
	TargetPeg = &([]int{})
	HelperPeg = &([]int{})
	return
}

func getNsortedNum(num int) *[]int {
	ser := []int{num}
	for i := 1; i < num; i++ {
		ser = append(ser, (num - i))
	}
	return &ser
}

func Print() {
	fmt.Print("Source >>>")
	fmt.Println(*(SourcePeg))
	fmt.Print("Helper >>>")
	fmt.Println(*(HelperPeg))
	fmt.Print("Target >>>")
	fmt.Println(*(TargetPeg))
}

var recur float64
var count float64

func SimulateGame() {
	recur = math.Exp2(float64(len(*(SourcePeg)))) - 1
	fmt.Println("Order of the algorithm" + strconv.Itoa(int(recur)))
	hanoi(len(*(SourcePeg)), SourcePeg, HelperPeg, TargetPeg)
	fmt.Println("Total moves: ", count)
}

func hanoi(n int, source *[]int, helper *[]int, target *[]int) {
	if n > 0 {
		recur = recur - 1
		count++
		//fmt.Println("**************" + strconv.Itoa(int(recur)))

		//		# move tower of size n - 1 to helper:
		//		hanoi(n - 1, source, target, helper)
		//		# move disk from source peg to target peg
		//		if source[0]:
		//		disk = source[0].pop()
		//		print "moving " + str(disk) + " from " + source[1] + " to " + target[1]
		//		target[0].append(disk)
		//		# move tower of size n-1 from helper to target
		//		hanoi(n - 1, helper, source, target)
		hanoi(n-1, source, target, helper)
		if len(*source) > 0 {
			disk := (*source)[len(*source)-1]
			*source = (*source)[:(len(*source) - 1)]
			*target = append((*target), disk)
			fmt.Println("target")
			fmt.Println(*target)
		}
		hanoi(n-1, helper, source, target)
	}
}
