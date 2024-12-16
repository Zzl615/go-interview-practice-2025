package main

import (
	"fmt"
)

// 浮点数的运算和整型一样，也要保证操作数的类型一致，float32 和 float64 类型数据不能混合运算，需要手动进行强制转化才可以，这一点和动态语言不同。
func typeConversion() {
	var floatValue1 float32 = 10.0
	floatValue2 := float64(10.0) // 如果不加小数点，floatValue2 会被推导为整型而不是浮点型
	floatValue3 := 1.1e-10
	fmt.Println(floatValue1, floatValue2, floatValue3)
}

// 浮点数精度问题
// 浮点数不是一种精确的表达方式，因为二进制无法精确表示所有十进制小数。
// 计算机底层将十进制的 0.1 和 0.7 转化为二进制表示时，会丢失精度，所以永远不要相信浮点数结果精确到了最后一位，也永远不要比较两个浮点数是否相等。
func ImpreciseCase() {
	floatValue4 := 0.1
	// 二进制表示
	fmt.Printf("%b\n", floatValue4)
	floatValue5 := 0.7
	// 二进制表示
	fmt.Printf("%b\n", floatValue5)
	floatValue6 := floatValue4 + floatValue5
	// 二进制表示
	fmt.Printf("%b\n", floatValue6)
	fmt.Println(floatValue6)
}

func main() {
	ImpreciseCase()
}
