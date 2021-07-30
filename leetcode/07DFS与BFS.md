#### [433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

难度中等80

一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 `"A"`, `"C"`, `"G"`, `"T"`中的任意一个。

假设我们要调查一个基因序列的变化。**一次**基因变化意味着这个基因序列中的**一个**字符发生了变化。

例如，基因序列由`"AACCGGTT"` 变化至 `"AACCGGTA" `即发生了一次基因变化。

与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。

现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

**注意：**

1. 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
2. 如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
3. 假定起始基因序列与目标基因序列是不一样的。

**示例 1：**

```
start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]

返回值: 1
```

**示例 2：**

```
start: "AACCGGTT"
end:   "AAACGGTA"
bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]

返回值: 2
```

**示例 3：**

```
start: "AAAAACCC"
end:   "AACCCCCC"
bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]

返回值: 3
```

```
func minMutation(start string, end string, bank []string) int {
    cList := []string{"A", "C", "G", "T"}
    visited := map[string]bool{}
    bSet := map[string]bool{}
    for _, b := range bank {
        visited[b] = false
        bSet[b] = true
    }
    visited[start] = false
    ans := 0
    queue := []string{start}
    for len(queue) != 0 {
        length := len(queue)
        ans++
        for l := 0; l < length; l++ {
            b := queue[0]
            queue = queue[1:]
            for i := 0; i < len(start); i++ {
                for j := 0; j < len(cList); j++ {
                    newBank := b[:i] + cList[j] + b[i + 1:]
                    if i, ok := bSet[newBank]; (ok && i == false) || ! ok { continue }
                    if visited[newBank] == true { continue }
                    if newBank == end { return ans }
                    visited[newBank] = true
                    queue = append(queue, newBank)
                }
            }
        }
        
    }
    return -1
}
```



#### [515. 在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)

难度中等135

您需要在二叉树的每一行中找到最大的值。

**示例：**

```
输入: 

          1
         / \
        3   2
       / \   \  
      5   3   9 

输出: [1, 3, 9]
```

```
func largestValues(root *TreeNode) []int {
    ans := []int{}
    if root == nil { return ans }
    
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        length := len(queue)
        temp := math.MinInt32
        for i := 0; i < length; i++ {
            node := queue[0]
            queue = queue[1:]
            temp = max(temp, node.Val)
            if node.Left != nil { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        ans = append(ans, temp)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```



#### [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)

难度困难794

字典 `wordList` 中从单词 `beginWord` 和 `endWord` 的 **转换序列** 是一个按下述规格形成的序列：

- 序列中第一个单词是 `beginWord` 。
- 序列中最后一个单词是 `endWord` 。
- 每次转换只能改变一个字母。
- 转换过程中的中间单词必须是字典 `wordList` 中的单词。

给你两个单词 `beginWord` 和 `endWord` 和一个字典 `wordList` ，找到从 `beginWord` 到 `endWord` 的 **最短转换序列** 中的 **单词数目** 。如果不存在这样的转换序列，返回 0。

**示例 1：**

```
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
```

**示例 2：**

```
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。
```

**提示：**

- `1 <= beginWord.length <= 10`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`、`endWord` 和 `wordList[i]` 由小写英文字母组成
- `beginWord != endWord`
- `wordList` 中的所有字符串 **互不相同**

```
func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := map[string]bool{}
	for i := 0; i < len(wordList); i++ {
		visited[wordList[i]] = true
	}

	queue := []string{beginWord}
	level := 1
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			for c := 0; c < len(word); c++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := word[:c] + string(j) + word[c + 1:]
					if visited[newWord] {
						queue = append(queue, newWord)
						visited[newWord] = false
					}
				}
			}
		}
		level++
	}
	return 0
}
```



#### [126. 单词接龙 II](https://leetcode-cn.com/problems/word-ladder-ii/)

难度困难447

按字典 `wordList` 完成从单词 `beginWord` 到单词 `endWord` 转化，一个表示此过程的 **转换序列** 是形式上像 `beginWord -> s1 -> s2 -> ... -> sk` 这样的单词序列，并满足：

- 每对相邻的单词之间仅有单个字母不同。
- 转换过程中的每个单词 `si`（`1 <= i <= k`）必须是字典 `wordList` 中的单词。注意，`beginWord` 不必是字典 `wordList` 中的单词。
- `sk == endWord`

给你两个单词 `beginWord` 和 `endWord` ，以及一个字典 `wordList` 。请你找出并返回所有从 `beginWord` 到 `endWord` 的 **最短转换序列** ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 `[beginWord, s1, s2, ..., sk]` 的形式返回。

**示例 1：**

```
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
解释：存在 2 种最短的转换序列：
"hit" -> "hot" -> "dot" -> "dog" -> "cog"
"hit" -> "hot" -> "lot" -> "log" -> "cog"
```

**示例 2：**

```
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：[]
解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
```

**提示：**

- `1 <= beginWord.length <= 7`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`、`endWord` 和 `wordList[i]` 由小写英文字母组成
- `beginWord != endWord`
- `wordList` 中的所有单词 **互不相同**

```
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    wordSet := map[string]bool{}
    visited := map[string]bool{}
    levelMap := map[string]int{}
    for i := 0; i < len(wordList); i++ {
        visited[wordList[i]] = false
        wordSet[wordList[i]] = true
        levelMap[wordList[i]] = -1
    }
    results := map[string][]string{}

    queue := []string{beginWord}
    visited[beginWord] = false
    level := -1
    for len(queue) != 0 {
        length := len(queue)
        level++
        for l := 0; l < length; l++ {
            word := queue[0]
            queue = queue[1:]
            temp := []string{}
            for i := 0; i < len(word); i++ {
                for c := 'a'; c <= 'z'; c++ {
                    newWord := word[:i] + string(c) + word[i + 1:]
                    if _, ok := wordSet[newWord]; !ok { continue }
                    if newWord == word { continue }
                    temp = append(temp, newWord)
                    if w, ok := levelMap[newWord]; ok && w > -1 { continue }
                    levelMap[newWord] = level
                    if visited[newWord] == true { continue }
                    visited[newWord] = true
                    queue = append(queue, newWord)
                }
            }
            results[word] = temp
        }
    }

    ans := [][]string{}
    var dfs func(temp []string, word string, level int)
    dfs = func(temp []string, word string, level int) {
        if word == endWord {
            temp = append(temp, endWord)
            t := make([]string, len(temp))
            copy(t, temp)
            ans = append(ans, t)
        }

        temp = append(temp, word)
        child := results[word]
        for i := 0; i < len(child); i++ {
            if levelMap[child[i]] == level {
                dfs(temp, child[i], level + 1)
            }
        }
        temp = temp[:len(temp) - 1]
    }
    
    dfs([]string{}, beginWord, 0)
    return ans
}



```



#### [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

难度中等1216

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**示例 1：**

```
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
```

**示例 2：**

```
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
```

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` 的值为 `'0'` 或 `'1'`

```
func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 { return 0 }
	m := len(grid[0])
	if m == 0 { return 0 }
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}

		}
	}
	return count
}

func dfs(grid [][]byte, i, j int)  {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i + 1, j)
	dfs(grid, i - 1, j)
	dfs(grid, i, j + 1)
	dfs(grid, i, j - 1)
}
```



#### [529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)

难度中等235

让我们一起来玩扫雷游戏！

给定一个代表游戏板的二维字符矩阵。 **'M'** 代表一个**未挖出的**地雷，**'E'** 代表一个**未挖出的**空方块，**'B'** 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的**已挖出的**空白方块，**数字**（'1' 到 '8'）表示有多少地雷与这块**已挖出的**方块相邻，**'X'** 则表示一个**已挖出的**地雷。

现在给出在所有**未挖出的**方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

1. 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 **'X'**。
2. 如果一个**没有相邻地雷**的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的**未挖出**方块都应该被递归地揭露。
3. 如果一个**至少与一个地雷相邻**的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
4. 如果在此次点击中，若无更多方块可被揭露，则返回面板。

**示例 1：**

```
输入: 

[['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'M', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E']]

Click : [3,0]

输出: 

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

解释:
```

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/minesweeper_example_1.png)

**示例 2：**

```
输入: 

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

Click : [1,2]

输出: 

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'X', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

解释:
```

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/minesweeper_example_2.png)

**注意：**

1. 输入矩阵的宽和高的范围为 [1,50]。
2. 点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
3. 输入面板不会是游戏结束的状态（即有地雷已被挖出）。
4. 简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。

```
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	dfs(board, click[0], click[1])
	return board
}

var dirX = []int{-1, -1, -1,  0, 0,  1, 1, 1}
var dirY = []int{-1,  0,  1, -1, 1, -1, 0, 1}

func dfs(board [][]byte, x, y int) {
	cnt := 0
	for i := 0; i < 8; i++ {
		tx := x + dirX[i]
		ty := y + dirY[i]
		if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			cnt++
		}
	}

	if cnt > 0 {
		board[x][y] = '0' + byte(cnt)
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx := x + dirX[i]
			ty := y + dirY[i]
			if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) {
				continue
			}
			if board[tx][ty] == 'E' {
				dfs(board, tx, ty)
			}
		}
	}
}
```



#### [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)

难度中等542

请你判断一个 `9x9` 的数独是否有效。只需要 **根据以下规则** ，验证已经填入的数字是否有效即可。

1. 数字 `1-9` 在每一行只能出现一次。
2. 数字 `1-9` 在每一列只能出现一次。
3. 数字 `1-9` 在每一个以粗实线分隔的 `3x3` 宫内只能出现一次。（请参考示例图）

数独部分空格内已填入了数字，空白格用 `'.'` 表示。

**注意：**

- 一个有效的数独（部分已被填充）不一定是可解的。
- 只需要根据以上规则，验证已经填入的数字是否有效即可。

**示例 1：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/04/12/250px-sudoku-by-l2g-20050714svg.png)

```
输入：board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：true
```

**示例 2：**

```
输入：board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
```

**提示：**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` 是一位数字或者 `'.'`

```
func isValidSudoku(board [][]byte) bool {
    rows := [9][9]bool{}
    cols := [9][9]bool{}
    boxes := [9][9]bool{}

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' { continue }
            idx := int(board[i][j] - '1')
            boxIdx := (i / 3) * 3 + j / 3
            if rows[i][idx] || cols[j][idx] || boxes[boxIdx][idx] { return false }
            rows[i][idx] = true
            cols[j][idx] = true
            boxes[boxIdx][idx] = true
        }
    }
    return true
}
```



#### [37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/)

难度困难876

编写一个程序，通过填充空格来解决数独问题。

数独的解法需 **遵循如下规则**：

1. 数字 `1-9` 在每一行只能出现一次。
2. 数字 `1-9` 在每一列只能出现一次。
3. 数字 `1-9` 在每一个以粗实线分隔的 `3x3` 宫内只能出现一次。（请参考示例图）

数独部分空格内已填入了数字，空白格用 `'.'` 表示。

**示例：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/04/12/250px-sudoku-by-l2g-20050714svg.png)

```
输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
```

**提示：**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` 是一位数字或者 `'.'`
- 题目数据 **保证** 输入数独仅有一个解

```
func solveSudoku(board [][]byte)  {
    var isvilad func(i, j int, val byte) bool
    isvilad = func(x, y int, val byte) bool {
        for i := 0; i < 9; i++ {
            if board[x][i] == val || board[i][y] == val { return false }
        }
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                tx := (x / 3) * 3
                ty := (y / 3) * 3
                if board[i + tx][j + ty] == val { return false }
            }
        }
        return true
    }
    var dfs func() bool
    dfs = func() bool {
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                if board[i][j] != '.' { continue }
                var k byte
                for k = '1'; k <= '9'; k++ {
                    if !isvilad(i, j, k) { continue }
                    board[i][j] = k
                    if dfs() {
                        return true
                    }
                    board[i][j]='.'
                }
                return false
            }
        }
        return true
    }
    dfs()
}
```



#### [1091. 二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)

难度中等106

给你一个 `n x n` 的二进制矩阵 `grid` 中，返回矩阵中最短 **畅通路径** 的长度。如果不存在这样的路径，返回 `-1` 。

二进制矩阵中的 畅通路径 是一条从 **左上角** 单元格（即，`(0, 0)`）到 右下角 单元格（即，`(n - 1, n - 1)`）的路径，该路径同时满足下述要求：

- 路径途经的所有单元格都的值都是 `0` 。
- 路径中所有相邻的单元格应当在 **8 个方向之一** 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。

**畅通路径的长度** 是该路径途经的单元格总数。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/02/18/example1_1.png)

```
输入：grid = [[0,1],[1,0]]
输出：2
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/02/18/example2_1.png)

```
输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
输出：4
```

**示例 3：**

```
输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
输出：-1
```

**提示：**

- `n == grid.length`
- `n == grid[i].length`
- `1 <= n <= 100`
- `grid[i][j]` 为 `0` 或 `1`

```
func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	if grid == nil || rows == 0 || grid[0][0] == 1 || grid[rows-1][rows-1] == 1 {
		return -1
	}
	if len(grid) == 1 && grid[0][0] == 0 {
		return 1
	}
	direction := []int{-1, 0, 1}
	grid[0][0] = 1
	//途经的每一个点都记录从起点到次的长度
	que := make([]int, 0)
	que = append(que, 0)
	//用que记录当前点的坐标，判断有没有下一个节点

	var x, y, nx, ny int

	for len(que) > 0 {
		x, y = que[0]/rows, que[0]%rows
		que = que[1:]
		for _, i := range direction {
			for _, j := range direction {
				if i == j && i == 0 {
					continue
				}
				nx, ny = x+i, y+j
				if nx < rows && ny < rows && nx >= 0 && ny >= 0 && grid[nx][ny] == 0 {
					que = append(que, nx*rows+ny)
					grid[nx][ny] = grid[x][y] + 1
					if nx == rows-1 && ny == rows-1 {
						return grid[nx][ny]
					}
				}
			}
		}
	}
	return -1
}
```



#### [733. 图像渲染](https://leetcode-cn.com/problems/flood-fill/)

难度简单191

有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标 `(sr, sc)` 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 `newColor`，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

**示例 1:**

```
输入: 
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析: 
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
```

**注意:**

- `image` 和 `image[0]` 的长度在范围 `[1, 50]` 内。
- 给出的初始点将满足 `0 <= sr < image.length` 和 `0 <= sc < image[0].length`。
- `image[i][j]` 和 `newColor` 表示的颜色值在范围 `[0, 65535]`内。

```
var (
    dx = []int{1, 0, 0, -1}
    dy = []int{0, 1, -1, 0}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    currColor := image[sr][sc]
    if currColor == newColor {
        return image
    }
    n, m := len(image), len(image[0])
    queue := [][]int{}
    queue = append(queue, []int{sr, sc})
    image[sr][sc] = newColor
    for i := 0; i < len(queue); i++ {
        cell := queue[i]
        for j := 0; j < 4; j++ {
            mx, my := cell[0] + dx[j], cell[1] + dy[j]
            if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == currColor {
                queue = append(queue, []int{mx, my})
                image[mx][my] = newColor
            }
        }
    }
    return image
}
```

```
var (
    dx = []int{1, 0, 0, -1}
    dy = []int{0, 1, -1, 0}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    currColor := image[sr][sc]
    if currColor != newColor {
        dfs(image, sr, sc, currColor, newColor)
    }
    return image
}

func dfs(image [][]int, x, y, color, newColor int) {
    if image[x][y] == color {
        image[x][y] = newColor
        for i := 0; i < 4; i++ {
            mx, my := x + dx[i], y + dy[i]
            if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
                dfs(image, mx, my, color, newColor)
            }
        }
    }
}
```

