// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	var cond sync.Cond
// 	cond.L = new(sync.Mutex)

// 	msgCh := make(chan int, 5)

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()
// 	rand.Seed(time.Now().UnixNano())

// 	producer := func(ctx context.Context, out chan<- int, idx int) {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				cond.Broadcast()
// 				fmt.Println("producer finished")
// 				return
// 			default:

// 				cond.L.Lock()

// 				for len(msgCh) == 5 {
// 					cond.Wait()
// 				}

// 				num := rand.Intn(500)
// 				out <- num
// 				fmt.Printf("producer: %d, msg: %d\n", idx, num)

// 				cond.Signal()

// 				cond.L.Unlock()
// 			}
// 		}
// 	}

// 	consumer := func(ctx context.Context, in <-chan int, idx int) {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				for len(msgCh) > 0 {
// 					select {
// 					case num := <-in:
// 						fmt.Printf("consumer %d, msg: %d\n", idx, num)
// 					default:
// 						break
// 					}
// 				}
// 				fmt.Println("consumer finished")
// 				return
// 			default:
// 				cond.L.Lock()

// 				for len(msgCh) == 0 {
// 					cond.Wait()
// 				}

// 				num := <-in
// 				fmt.Printf("consumer %d, msg: %d\n", idx, num)

// 				cond.Signal()

// 				cond.L.Unlock()
// 			}
// 		}
// 	}

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go producer(ctx, msgCh, i+1)
// 	}
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go consumer(ctx, msgCh, i+1)
// 	}

// 	wg.Wait()
// 	close(msgCh)
// 	fmt.Println("all finished")
// }

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	msgCh := make(chan int, 5)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rand.Seed(time.Now().UnixNano())

	producer := func(ctx context.Context, out chan<- int, idx int) {
		defer wg.Done()
		for {
			num := rand.Intn(500)

			select {
			case <-ctx.Done():
				fmt.Println("producer finished")
				return
			case out <- num:
				fmt.Printf("producer: %d, msg: %d\n", idx, num)
			}
		}
	}

	consumer := func(ctx context.Context, in <-chan int, idx int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				for {
					select {
					case num := <-in:
						fmt.Printf("consumer %d, msg: %d\n", idx, num)
					default:
						fmt.Println("consumer finished")
						return
					}
				}
			case num := <-in:
				fmt.Printf("consumer %d, msg: %d\n", idx, num)
			}
		}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go producer(ctx, msgCh, i+1)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go consumer(ctx, msgCh, i+1)
	}

	wg.Wait()
	close(msgCh)
	fmt.Println("all finished")
}
