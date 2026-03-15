package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			<-ch1 // 消费
			fmt.Println(i, "a")
			ch2 <- struct{}{}
		}
		<-ch1
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println(i, "b\n")
			ch1 <- struct{}{}
		}
	}()

	// 让ch1获取信号，开始打印
	ch1 <- struct{}{}
	wg.Wait()
	close(ch1)
	close(ch2)

	fmt.Println("finish")
}
