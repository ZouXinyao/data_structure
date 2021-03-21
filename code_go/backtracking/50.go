package backtracking

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	res := 1.0
	temp := x
	for n > 0 {
		if n % 2 == 1 {
			res *= temp
		}
		temp *= temp
		n /= 2
	}
	return res
}

