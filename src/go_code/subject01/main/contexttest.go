package main

import (
	"context"
	"fmt"
	"time"
)

//  communicating sequential processes (CSP) and software transactional memory (STM).
//并发模型，通信顺序过程，软件事物内存
const DbAddress = "db_address"
const CalculateValue = "calculate value"

func readDB(ctx context.Context, cost time.Duration) {
	fmt.Println("db address is ,", ctx.Value(DbAddress))

	select {
	case <-time.After(cost): //数据库读取
		fmt.Println("read data from db")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //任务需取消的原因
	}

}

func calculate(ctx context.Context, cost time.Duration) {

	fmt.Println("calculate value is", ctx.Value(CalculateValue))
	select {
	case <-time.After(cost): //数据计算
		fmt.Println("calculate finish")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, DbAddress, "localhost:3306")
	ctx = context.WithValue(ctx, CalculateValue, 1234)
	//设置子context 2s 后超时执行返回
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	//设置执行时间为4s
	go readDB(ctx, time.Second*4)
	go calculate(ctx, time.Second*4)

	time.Sleep(time.Second * 5)

}
