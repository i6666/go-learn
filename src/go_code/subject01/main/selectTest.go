package main

import (
	"fmt"
	"time"
)

func sendMsg(ch chan int, begin int) {
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}

func receiver(ch <-chan int) {
	val := <-ch
	fmt.Println("receive:", val)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go sendMsg(ch1, 0)
	go receiver(ch2)

	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1:
			fmt.Printf("get value %d from ch1 \n", val)
		case ch2 <- 2: //使用ch2发送消息
			fmt.Println("send value by ch2")
		case <-time.After(2 * time.Second):
			fmt.Println("time out")
			return
		}
	}

}
