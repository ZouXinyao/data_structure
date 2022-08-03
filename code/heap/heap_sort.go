package main

import "fmt"

func buildHeap(a []int, n int) {
	for i := n/2; i >= 1; i-- {
		heapify(a, n, i)
	}
}

func heapify(a []int, n int, i int) {
	for {
		// 选出要交换元素的位置，子节点中较大的那个。
		maxIndex := i
		if i*2 <= n && a[i] < a[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= n && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i {
			break
		}

		// 交换，然后节点索引下移
		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}
}

func heapSort(a []int, n int)  {
	buildHeap(a, n)
	for i := n; i > 1; {
		a[i], a[1] = a[1], a[i]
		i--
		heapify(a, i, 1)
	}
}

func main() {
	a := []int{0, 9, 21, 20, 17, 11, 12, 6, 15, 13, 1, 8}
	heapSort(a, len(a) - 1)
	fmt.Println(a)
}