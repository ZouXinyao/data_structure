package list

func climbStairs(n int) int {
	if n <= 3 { return n }
	first, second := 1, 2
	res := 0
	for i := 3; i <= n; i++{
		res = first + second
		first, second = second, res
	}
	return res
}
