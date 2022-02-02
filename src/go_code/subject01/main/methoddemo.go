package main

import (
	"math"
)

//func main() {
//	v := Vertex{3, 4}
//	v.Scale(10)
//
//	fmt.Println(v.Abs())
//
//	fmt.Println(Abs(v))
//
//	f := MyFloat(-math.Sqrt2)
//	fmt.Println(f.Abs())
//
//
//	v2 := Vertex{3, 4}
//	Scale(&v2,10)
//	fmt.Println(Abs(v2))
//
//	v3 := Vertex{3,4}
//	v3.Scale(2)
//	ScaleFunc(&v3,10)
//
//	p :=Vertex{4,3}
//	//方法与指针重定向 及时v是值 非指针go 会转换为(&p).Scale(3)
//	p.Scale(3)
//	ScaleFunc(&p,8)
//
//	fmt.Println(v,p)
//
//
//
//}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Vertex struct {
	X, Y float64
}

// 方法就是带有特殊接受参数的函数，位于方法名和关键字之间
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接受者的方法，可以修改接受者指向的值，如果用值接受者，仅会对副本进行操作
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 接受者的类型定义和方法声明需要在同一个包内
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}
