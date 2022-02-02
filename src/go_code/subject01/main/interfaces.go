package main

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describer(i I) {
	fmt.Printf("(%v ,%T)\n", i, i)
}

type IPAddr [4]byte

func (a IPAddr) String() string {
	var vv string
	for i, v := range a {
		if i == 0 {
			vv = fmt.Sprint(v)
		} else {
			vv = fmt.Sprintf("%v.%v", vv, v)
		}
	}
	return vv
}

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.age)
}

func main() {
	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex{3,4}
	//
	//a = f
	//a = &v
	//fmt.Println(a.Abs())

	//var i I
	//i = &T{"hello"}
	//describer(i)
	//i.M()
	//
	//i = F(math.Pi)
	//
	//describer(i)
	//i.M()
	//

	var i interface{} = "hello"
	//类型断言
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	do(21)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("I don't know aboutt type %T !\n", v)
	}

	a := Person{"kangkang", 20}
	b := Person{"kangkang", 20}
	fmt.Println(a, b)

}
