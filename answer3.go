package main

import (
	"fmt"
	"container/list"
	"sort"
)

func main() {
	twoNum(10,2)

	test(4,3,2,1)

	arrayMap([]int{1, 2, 3}, func(i int) int { return i * 5 })

	findMaxAndMin([]int{1, 20, 13, 5, 7})

	arraySort([]int{1, 20, 13, 5, 7})

	doSayHi := sayHi("jojo")
	doSayHi()
}

// 輸入兩個數字, 依照大小輸出
func twoNum(a, b int) {
	if a > b {
		fmt.Println(a, b)
	} else {
		fmt.Println(b, a)
	}
}

// 存放整數queue, 並後進先出
func test(a ...int) {
	queue := list.New()
	for _, v := range a {
		queue.PushBack(v)
	}

	for queue.Len() > 0 {
		fmt.Println(queue.Remove(queue.Back()))
	}
}

// 實作php的array_map
func arrayMap(input []int, f func(int) int){
	for i, v := range input {
		input[i] = f(v)
	}

	fmt.Println(input)
}

// 找 slice 中的最大值、最小值
func findMaxAndMin(input []int) {
	max := input[0]
	min := input[0]

	for _, v := range input {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	fmt.Println("最大值為:", max, ", 最小值為:", min)
}

// 實作php的array_sort
func arraySort(input []int) {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	fmt.Println(input)
}

// 實作一個回傳閉包的方法
func sayHi(name string) func() {
	return func() {
		fmt.Println("Hi " + name)
	}
}