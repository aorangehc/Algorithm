package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建slice
	ss := make([]int, 100)

	for i := 0; i < 100; i++ {
		ss[i] = i
	}

	ch := make(chan struct{}, 10) // 缓冲10

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("val: %d \n", ss[idx])
			<-ch
		}(i)
	}

	wg.Wait()

	close(ch)

	fmt.Println("finish")
}
