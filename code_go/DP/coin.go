package main

import (
	"fmt"
	"math"
)

func coinChange01(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := math.MaxInt32
	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		tmp := coinChange01(coins, amount-coin)
		if tmp == -1 {
			continue
		}
		res = min(res, tmp+1)
	}
	return getRes(res)
}

func coinChange02(coins []int, amount int) int {
	tempList := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		tempList[i] = -2
	}
	return helper02(coins, amount, &tempList)
}

func helper02(coins []int, amount int, tempList *[]int) int {
	if amount == 0 {
		return 0
	}
	if (*tempList)[amount] != -2 {
		return (*tempList)[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		tmp := coinChange01(coins, amount-coin)
		if tmp == -1 {
			continue
		}
		res = min(res, tmp+1)
	}
	(*tempList)[amount] = getRes(res)
	return getRes(res)
}

func coinChange03(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount
	}
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	return getRes(dp[amount])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getRes(res int) int {
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func main() {
	coins := []int{1, 2, 5, 10}
	amount := 28
	fmt.Println(coinChange01(coins, amount))
	fmt.Println(coinChange02(coins, amount))
	fmt.Println(coinChange03(coins, amount))

}
