package main

import (
	"fmt"
	"strconv"
)

func hello() {
	fmt.Print("hello")
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum/2 + 3
	y = sum - x
	return
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func add(x, y int) int {
	return x + y
}

func strConv() {
	fmt.Println("天龙八部\t未完待续")
	fmt.Println("天龙八部\n未完待续")
	fmt.Println("天龙八部,第二\r未完待续") //\r 当前行的最前面开始输出，覆盖以前的内容

	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)

	fmt.Printf("b type %t  b value %v \n", b, b)
}

//go 的每一个文件都要属于一个包
//go 语言指针使用的特点
func testPtr(num *int) {
	*num = 2

}

//go 函数可以直接支持返回多个值
func getSumAndSun(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
