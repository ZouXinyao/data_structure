package main

import "fmt"

type Heap struct {
	a     []int
	n     int
	count int
}

//init heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		a:     make([]int, capacity+1),
		n:     capacity,
		count: 0,
	}
}

// 添加元素，添加到末尾，然后向上堆化
func (heap *Heap) insert(data int) {
	// 堆满了
	if heap.count >= heap.n {
		return
	}

	heap.count++
	heap.a[heap.count] = data // 堆尾添加元素

	// 向上对比交换
	i := heap.count
	for i/2 > 0 && heap.a[i/2] < heap.a[i] {
		heap.a[i], heap.a[i/2] = heap.a[i/2], heap.a[i]
		i = i / 2
	}
}

func (heap *Heap) remove() {

	// 堆中没数据
	if heap.count == 0 {
		return
	}

	// 末尾的数据移到堆顶
	heap.a[1] = heap.a[heap.count]
	heap.count--

	// 向下堆化
	heapifyUpToDown(heap.a, heap.count)
}

func heapifyUpToDown(a []int, count int) {
	for i := 1; i <= count/2; {
		// 选出要交换元素的位置，子节点中较大的那个。
		maxIndex := i
		if i*2 <= count && a[i] < a[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
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

func main() {
	// heap := NewHeap(10)
	// heap.insert(6)
	// heap.insert(3)
	// heap.insert(9)
	// heap.insert(2)
	// heap.insert(5)
	// heap.insert(4)
	// fmt.Println(heap.a[:heap.count + 1])
	// heap.remove()
	// heap.remove()
	// fmt.Println(heap.a[:heap.count + 1])
	fmt.Println("255" > "256")
}
