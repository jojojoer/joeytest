package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk 我是中文"

	printCount(str)
	replace(str)
	reverseString(str)
}

func printCount(str string) {
    fmt.Println("字數量為:", utf8.RuneCountInString(str))
    fmt.Println("Byte數為:", len(str))
}

func replace(str string) {
	str = strings.Replace(str, str[4:7], "abc", 1)
	fmt.Println(str)
}

func reverseString(str string) {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
}
