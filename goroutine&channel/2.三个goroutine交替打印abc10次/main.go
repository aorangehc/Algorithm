package main

import (
	"fmt"
	"sync"
)

func main() {
	// channel 作为信号的作用，收到信号就打印
	// 不然就阻塞
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(3)

	// 打印 a
	go func() {
		defer wg.Done() // 线程结束之后关闭
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("a")
			ch2 <- struct{}{}
		}
		<-ch1 //第十次的时候打印 c的goroutine 写入 ch1
		// 防止阻塞，消费一下 ch1
	}()

	// 打印 b
	go func() {
		defer wg.Done() // 线程结束之后关闭
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("b")
			ch3 <- struct{}{}
		}
	}()

	// 打印 c
	go func() {
		defer wg.Done() // 线程结束之后关闭
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println("c\n")
			ch1 <- struct{}{}
		}
	}()

	// 启动
	ch1 <- struct{}{}

	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)

	fmt.Println("finish")
}
