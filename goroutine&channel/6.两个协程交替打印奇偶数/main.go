package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	wg.Add(2)

	// 偶数
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				<-ch2
				fmt.Printf("偶数%d\n", i)
				ch1 <- struct{}{}
			}
		}
		// <-ch2
	}()

	// 奇数
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				<-ch1
				fmt.Printf("奇数：%d\n", i)
				ch2 <- struct{}{}
			}
		}
		<-ch1
	}()

	ch1 <- struct{}{}

	wg.Wait()
}
