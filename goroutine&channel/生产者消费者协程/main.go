/*
- 题目表述：创建一个worker（在main方法中），包含一个生产者和消费者，生产者和消费者通过协程实现，生产者和消费者之间通过channel进行通信
- 要求：worker在阻塞5分钟之后优雅退出，结束协程，关闭main
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 生产者函数，每秒生产一个递增整数，发送到通道
// ctx：上下文对象，监听取消信号
// ch：通道，发送数据
func producer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()                           // 通知WaitGroup，此协程已完成
	ticker := time.NewTicker(1 * time.Second) // 创建每秒出大的定时器
	defer ticker.Stop()                       // 函数退出的时候停止定时器

	i := 0

	for {
		select {
		case <-ctx.Done(): // 监听上下文取消信号，监听到就退出
			fmt.Println("生产者停止")
			return // 退出协程
		case <-ticker.C: // 每秒触发一次
			i += 1
			select {
			case ch <- i:
				fmt.Println("produced:", i)
			case <-ctx.Done():
				fmt.Println("生产者停止")
				return
			}
		}
	}
}

func consumer(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("消费者停止")
			return
		case data, ok := <-ch: // 从通道接收数据
			if !ok {
				fmt.Println("通道关闭且无数据")
				return // 退出协程
			}
			fmt.Println("consumed:", data)
		}
	}
}

func main() {
	ch := make(chan int, 10) // 创建带有10个缓冲槽的通道
	var wg sync.WaitGroup    // 协程管理

	// 创建阻塞5分钟的上下文，到达5分钟的时候调用cancel()，释放上下文资源
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel() // 之后调用ctx.Done()

	// 启动生产者和消费者协程
	wg.Add(2)
	go producer(ctx, ch, &wg)
	go consumer(ctx, ch, &wg)

	// 等待上下文超时，一共5分钟
	<-ctx.Done()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("主进程结束")
}
