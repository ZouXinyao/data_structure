package stack_and_queue

func isValid(s string) bool {
	l := []byte(s)
	stack := []byte{}
	for _, b := range l {
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		} else if b == ')' {
			if len(stack) == 0 || stack[len(stack) - 1] != '(' { return false }
			stack = stack[:len(stack) - 1]
		} else if b == ']' {
			if len(stack) == 0 || stack[len(stack) - 1] != '[' { return false }
			stack = stack[:len(stack) - 1]
		} else {
			if len(stack) == 0 || stack[len(stack) - 1] != '{' { return false }
			stack = stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0
}

