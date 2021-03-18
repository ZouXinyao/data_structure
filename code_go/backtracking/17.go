package backtracking

var phoneMap map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 { return []string{} }

	digitsByte := []byte(digits)
	var res  []string
	var temp []byte
	var helper func(int)
	// 这次idx可以看成是level了，也就是digits的第几位
	helper = func(idx int) {
		if len(temp) == len(digits) {
			t := make([]byte, len(temp))
			copy(t, temp)
			res = append(res, string(t))
			return
		}
		letters := []byte(phoneMap[digitsByte[idx]])
		// 因为要遍历letters中的所有元素，所以从0开始
		for i := 0; i < len(letters); i++ {
			temp = append(temp, letters[i])
			helper(idx + 1)
			temp = temp[:len(temp) - 1]
		}
	}

	helper(0)
	return res
}
