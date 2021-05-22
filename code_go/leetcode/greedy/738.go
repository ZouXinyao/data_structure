package main

// 可以用整型处理，也可以转换成字符串处理；
// 处理的过程都是从低位到高位遍历，当N[i+1]位>N[i]位时，就把N[i+1]-1，N[i]及低所有位都变成9，遍历所有位。
// 举例： 4321->4319->4299->3999
func monotoneIncreasingDigits(N int) int {
	res := 0
	seed := 1
	for N > 0 {
		low := N % 10
		N /= 10
		high := N % 10
		if high > low {
			res = seed * 10 - 1
			N--
		} else {
			res = low * seed + res
		}
		seed *= 10
	}
	return res
}

