package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//将sum 发送到 信道
	c <- sum
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	//s := [] int{7,2,8,-9,4,0}
	////信道使用前必须创建
	//c := make(chan int)
	//
	//go sum(s[:len(s)/2],c)
	//go sum(s[len(s)/2:],c)
	////从信道中接收
	//x,y := <-c ,<-c
	//fmt.Println(x,y,x+y)

	//c := make(chan int,10)
	//go fibonacci2(cap(c),c)
	//
	//for i := range c {
	//	fmt.Println(i)
	//}

	c := make(chan int)

	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci3(c, quit)

}
