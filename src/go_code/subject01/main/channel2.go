package main

import (
	"fmt"
	"time"
)

func Producer(begin, end int, queue chan<- int) {
	for i := begin; i < end; i++ {
		fmt.Println("produce: >>>>>>>>>>>>>>>", i)
		queue <- i
	}

}

func Consume(queue <-chan int) {
	for val := range queue {
		fmt.Println("consume:<<<<<<<<<<<<<<<", val)
	}
}

func send(ch chan int, begin int) {
	//循环向channel 发送消息
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}
func receive(ch <-chan int) {
	val := <-ch
	fmt.Println("receive:", val)
}

func main() {
	queue := make(chan int)

	defer close(queue)

	for i := 0; i < 3; i++ {
		//多个生产者 0,5  5 10 10 15
		go Producer(i*5, (i+1)*5, queue)

	}
	//单个消费者
	go Consume(queue)
	time.Sleep(time.Second)

}
