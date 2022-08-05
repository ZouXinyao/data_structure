#### [208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

**[Trie](https://baike.baidu.com/item/字典树/9825209?fr=aladdin)**（发音类似 "try"）或者说 **前缀树** 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

- `Trie()` 初始化前缀树对象。
- `void insert(String word)` 向前缀树中插入字符串 `word` 。
- `boolean search(String word)` 如果字符串 `word` 在前缀树中，返回 `true`（即，在检索之前已经插入）；否则，返回 `false` 。
- `boolean startsWith(String prefix)` 如果之前已经插入的字符串 `word` 的前缀之一为 `prefix` ，返回 `true` ；否则，返回 `false` 。

**示例：**

```
输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
```

**提示：**

- `1 <= word.length, prefix.length <= 2000`
- `word` 和 `prefix` 仅由小写英文字母组成
- `insert`、`search` 和 `startsWith` 调用次数 **总计** 不超过 `3 * 104` 次

```go
type Trie struct {
    subTrie map[byte]*Trie
    isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        subTrie: map[byte]*Trie{},
        isEnd: false,
    }
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	root := t
	for i := 0; i < len(word); i++ {
		node, ok := root.subTrie[word[i]]
		if ok {
			root = node
		} else {
			newNode := Constructor()
			root.subTrie[word[i]] = &newNode
			root = root.subTrie[word[i]]
		}
	}
	root.isEnd = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
    root := t
    for i := 0; i < len(word); i++ {
        node, ok := root.subTrie[word[i]]
        if ok {
            root = node
        } else {
            return false
        }
    }
    return root.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
    root := t
    for i := 0; i < len(prefix); i++ {
        node, ok := root.subTrie[prefix[i]]
        if ok {
            root = node
        } else {
            return false
        }
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
```

#### [212. 单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)

给定一个 `m x n` 二维字符网格 `board` 和一个单词（字符串）列表 `words`，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过 **相邻的单元格** 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/11/07/search1.jpg)

```
输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
输出：["eat","oath"]
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/11/07/search2.jpg)

```
输入：board = [["a","b"],["c","d"]], words = ["abcb"]
输出：[]
```

**提示：**

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` 是一个小写英文字母
- `1 <= words.length <= 3 * 104`
- `1 <= words[i].length <= 10`
- `words[i]` 由小写英文字母组成
- `words` 中的所有字符串互不相同

```go
// 按照模板写UnionFound也可以，不过回溯的条件不好弄，会超时。
func findWords(board [][]byte, words []string) []string {
	root := Constructor()
	for i := 0; i < len(words); i++ {
		root.Insert(words[i])
	}
	res := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, &res, &root)
		}
	}
	return res
}

var dirX = []int{0, 1, -1, 0}
var dirY = []int{-1, 0, 0, 1}

func dfs(board [][]byte, x, y int, res *[]string, node *Trie) {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]){
		return
	}
	temp := board[x][y]
	
	if _, ok := node.subTrie[temp]; !ok {
		return
	}
	node = node.subTrie[temp]
	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}


	board[x][y] = '*'
	for i := 0; i < 4; i++ {
		tx := dirX[i] + x
		ty := dirY[i] + y
		dfs(board, tx, ty, res, node)
	}
	board[x][y] = temp
}

type Trie struct {
	subTrie map[byte]*Trie
	word string
}


func Constructor() Trie {
	return Trie{
		subTrie: map[byte]*Trie{},
		word: "", // 标识Trie中存在的word，方便删除。
	}
}

func (this *Trie) Insert(word string)  {
	root := this
	for i := 0; i < len(word); i++ {
		node, ok := root.subTrie[word[i]]
		if ok {
			root = node
		} else {
			newNode := Constructor()
			root.subTrie[word[i]] = &newNode
			root = root.subTrie[word[i]]
		}
	}
	root.word = word
}

```





