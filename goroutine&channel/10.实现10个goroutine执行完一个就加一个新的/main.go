package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 10)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		ch <- struct{}{}

		go func(id int) {
			defer wg.Done()
			fmt.Printf("id: %d", id)
			<-ch
		}(i)
	}

	wg.Wait()
}
