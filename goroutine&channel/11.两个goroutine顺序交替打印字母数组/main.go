package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 'a'; i <= 'z'; i++ {
			fmt.Print(string(i))
			ch <- struct{}{}
			<-ch
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-ch
			fmt.Println(i)
			ch <- struct{}{}
		}
	}()

	wg.Wait()

	close(ch)
}
