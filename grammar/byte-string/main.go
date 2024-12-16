package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 字面值，表示泰语中的单词 “hello”
	// 字面值是字符串的直接表示，通常用于初始化字符串变量。
	// 字面值可以包含任何有效的 UTF-8 编码的字符，包括非 ASCII 字符。
	// 在编程中，字面值（literal value）指的是在源代码中直接表示其自身值的数据。
	// 这意味着该值是固定的，不是由变量或计算结果得来的。
	// 字面值可以是数字、字符、字符串或其他固定格式的数据。
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
