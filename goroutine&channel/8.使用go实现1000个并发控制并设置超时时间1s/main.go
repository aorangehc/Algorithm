package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1000)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)

	defer cancel()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		ch <- i
		go func(id int) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				fmt.Printf("协程 %d 被取消或超时退出\n", id)
				return
			case val := <-ch:
				// 正常从通道读取数据并处理
				fmt.Printf("协程 %d 成功处理数据: %d\n", id, val)

			}
		}(i)
	}

	wg.Wait()

	close(ch)

	fmt.Println("finish")
}
