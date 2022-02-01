package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"strings"
)

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	//将函数推迟至外层函数 返回后执行
	defer fmt.Println("end")
	//strConv()
	//fmt.Println("my favorite num is",rand.Intn(10))
	//fmt.Printf("now you have %g problems", math.Sqrt(16))
	//fmt.Println(math.Pi)
	//
	//fmt.Println(add(10,12))
	// 简洁赋值语句，可在类型明确的地方代理var 声明，不能在函数外使用
	//a, b := swap("好", "你")
	//fmt.Print(a,b)
	//
	//fmt.Println(split(20))

	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value :%v \n", Tobe, Tobe)
	fmt.Printf("Type: %T Value :%v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value :%v \n", z, z)

	var i int = 42
	//go 的不同类型需要显示转换
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("Type: %T Value :%v \n", u, u)

	//testFor()
	//testWhile()
	//testSwitch()

	//testStruct()

	//testArray()

	//testSlice()

	//testSlice2()

	//testSlice3()

	//testMakeSlice()

	//testAppend()

	//var pow =  []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for i,v:= range pow{
	//	fmt.Printf("2**%d = %d \n" ,i,v)
	//}
	//testMake()
	//testMap1()
	//testMap2()
	//s := "I am learning Go Go !"
	//
	//var  result  map[string]int= wordCount(s)
	//fmt.Println(result)
	//testFun()

	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}
	//
	n := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(n())
	}
}

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		j := x
		x = y
		y = j + y
		return j
	}
}

func testFun() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func wordCount(s string) map[string]int {
	ss := strings.Split(s, " ")
	m := make(map[string]int)

	for _, si := range ss {
		v, ok := m[si]
		if ok {
			v++
		} else {
			v = 1
		}
		m[si] = v
	}
	return m

}

func testMap2() {
	m := make(map[string]int)
	m["kangkang"] = 22
	fmt.Println("the value", m["kangkang"])

	m["kangkang"] = 50
	fmt.Println("the value", m["kangkang"])

	delete(m, "kangkang")

	fmt.Println("delete the value", m["kangkang"])

	//通过双赋值检测key 是否存在，不存在ok 返回false,否则为true
	v, ok := m["kangkang"]
	fmt.Println("ok = ", ok, "v = ", v)
}

func testMap1() {
	m := make(map[string]Student)
	m["kkk"] = Student{"kangkang", 23}

	var mm = map[string]Student{
		"kkk":   {"kangkang", 20},
		"ddkkk": {"kangkang", 20},
	}

	fmt.Println(m)
	fmt.Println(mm)
}

func testMake() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2 ** i
	}
	for j, _ := range pow {
		fmt.Printf("j = %d \n", j)
	}

	for _, v := range pow {
		fmt.Printf("v= %d \n", v)
	}
}

func testAppend() {
	board := [][]string{[]string{"_", "_", "_"}, []string{"_", "_", "_"}, []string{"_", "_", "_"}}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func testMakeSlice() {
	a := make([]int, 5) //len=5 cap=5 [0 0 0 0 0]
	printSlice(a)

	b := make([]int, 0, 5) //len=0 cap=5 []
	printSlice(b)

	c := b[:2]
	printSlice(c) //len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice(d) //len=3 cap=3 [0 0 0]
}

func testSlice3() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(ints)
	//长度0
	ints = ints[:0]
	printSlice(ints)
	//扩展长度
	ints = ints[:4]
	printSlice(ints)
	//舍弃前两个
	ints = ints[2:]
	printSlice(ints)

	var n []int
	printSlice(n)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}

func testSlice2() {
	bools := [3]bool{true, false, false}
	i2 := []bool{true, false, false}
	p1 := &bools
	p2 := &i2

	fmt.Println(p1, p2)
}

func testSlice() {
	nums := [6]int{1, 2, 3, 4, 5, 6}
	//排除最后一个元素，【1，2，3】
	//切片不会存储任何数据，只是描述底层数组中的一段
	s := nums[0:3]
	fmt.Println(s)
}

func testArray() {
	var a [10]int
	a[0] = 2
	a[1] = 2

	b := [2]string{"hello"}
	fmt.Println(a, b)
}

func testStruct() {
	var p *int
	n := 10
	p = &n
	fmt.Println(p, *p)
	s := Student{"kangkang", 25}

	pp := &s
	pp.name = "kk"
	fmt.Println(*pp)
}

//没有条件的switch 和 switch true 一样，switch{} ，比if-then-else 清晰
func testSwitch() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux 系统")
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Println(os, "系统")

	}

}

type Student struct {
	name string
	age  int
}

func testFor() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	return sum

}

func testWhile() {
	sum := 0
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

}

var i, j int = 1, 2

var c, python, java = true, false, "no"
