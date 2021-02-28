package main

func lemonadeChange(bills []int) bool {
	money := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}

	for _, bill := range bills {
		money[bill]++

		if bill == 10 {
			money[5]--
			if money[5] < 0 {
				return false
			}
		} else if bill == 20 {
			money[5]--
			if money[10] == 0 {
				money[5] -= 2
			} else {
				money[10]--
			}
			if money[5] < 0 || money[10] < 0 {
				return false
			}
		}
	}
	return true
}