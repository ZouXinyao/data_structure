package main

func solveSudoku(board [][]byte)  {
	// 判断是否能把这个数放进数组
	var isValid func(row, col int, val byte) bool
	isValid = func(row, col int, val byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == val { return false }
		}
		for i := 0; i < 9; i++ {
			if board[i][col] == val { return false }
		}

		startRow := (row / 3) * 3;
		startCol := (col / 3) * 3;
		for i := startRow; i < startRow + 3; i++ {
			for j := startCol; j < startCol + 3; j++ {
				if board[i][j] == val { return false }
			}
		}
		return true
	}
	// 这里要有一个返回值，因为if helper() { return true }这里要把k保存在[i][j]位置
	var helper func() bool
	helper = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' { continue }
				var k byte
				for k = '1'; k <= '9'; k++ {
					if isValid(i, j, k) {
						board[i][j] = k
						if helper() { return true }// 如果找到合适一组立刻返回
						board[i][j] = '.'
					}
				}
				// 没有合适的就返回false
				return false
			}
		}
		return true
	}
	helper()
}