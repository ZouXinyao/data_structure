package main

import (
	"sort"
)

func largestSumAfterKNegations(ary []int, k int) int {
	sort.Ints(ary)
	for i, mID := 0, 0; i < k; i++ {
		ary[mID] = -ary[mID]
		if mID+1 < len(ary) && ary[mID] > ary[mID+1] {
			mID++
		}
	}
	sum := 0
	for _, a := range ary {
		sum += a
	}
	return sum
}