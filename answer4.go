package main

import (
	"fmt"
)

func main() {
	ints := []int{1, 2, 3}
	arrayMapInterface(ints, func(i interface{}) interface{} {
		return i.(int) * 2
	})

	strings := []string{"hihi", "hey", "hello"}
	arrayMapInterface(strings, func(i interface{}) interface{} {
		return i.(string) + "!!"
	})
}

// 實作php的array_map
func arrayMapInterface(input interface{}, f func(interface{}) interface{}) {
	switch input.(type) {
	case []int:
		for i, v := range input.([]int) {
			input.([]int)[i] = f(v).(int)
		}
		fmt.Println(input)
	case []string:
		for i, v := range input.([]string) {
			input.([]string)[i] = f(v).(string)
		}
		fmt.Println(input)
	default:
		fmt.Println("Unsupported type")
	}
}

