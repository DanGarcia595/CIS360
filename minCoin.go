package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{1, 5, 10, 25}
	fmt.Println(array)
	fmt.Println(mincoin(78, array))

}

func maxsum(a []int) int {
	solution := make([]int, len(a)+1)

	solution[0] = 0

	for j := 1; j < len(solution); j++ {
		solution[j] = int(math.Max(float64(solution[j-1]+a[j-1]), float64(a[j-1])))
	}

	result := solution[0]
	for j := 1; j < len(solution); j++ {
		if result < solution[j] {
			result = solution[j]
		}
	}

	return result

}

func mincoin(amount int, coins []int) int {
	fmt.Println("Making change for ", amount, "cents")
	coinReq := make([]int, amount+1)
	coinarray := make([]int, 0)

	CC := make([]int, len(coins))

	coinReq[0] = 0

	for amt := 1; amt <= amount; amt++ {
		for j := 0; j < len(CC); j++ {
			CC[j] = -1
		}

		for j := 0; j < len(coins); j++ {
			if coins[j] <= amt {
				CC[j] = coinReq[amt-coins[j]] + 1
			}
		}
		coinReq[amt] = -1
		for j := 0; j < len(CC); j++ {
			if CC[j] > 0 && (coinReq[amt] == -1 || coinReq[amt] > CC[j]) {
				coinReq[amt] = CC[j]
				//coinarray = append(coinarray, CC[j])
				//coinarray[i] = coins[CC[j]]
				//i++

			}
		}
	}
	j := 0
	temp := amount
	for {
		if coinReq[temp]-1 == coinReq[temp-coins[j]] {
			coinarray = append(coinarray, coins[j])
			temp -= coins[j]
		} else {
			j++
		}
		if len(coinarray) >= coinReq[amount] {
			break
		}
	}
	fmt.Println(coinarray)
	return coinReq[amount]

}
