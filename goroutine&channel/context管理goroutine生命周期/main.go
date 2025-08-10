package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个可取消的context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 确保函数结束时取消context

	// 启动一个goroutine，监听cancel信号
	go func(ctx context.Context) {
		// 判断是否有设置Deadline，如果没有则为设置了cancel函数
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Println("Deadline set for:", deadline)
		} else {
			fmt.Println("No deadline set.")
		}
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine被取消")
				return
			default:
				fmt.Println("Goroutine运行中")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	// 运行3秒后取消Context
	<-ctx.Done()

	// 等待一会儿以便观察输出
	time.Sleep(1 * time.Second)
	fmt.Println("主程序结束")
}
