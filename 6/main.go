package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Первый способ через контекст
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutinde 1 is done")
				return
			default:
				fmt.Println("loop 1")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()

	// Второй способ через сигнальный канал
	exitChannel := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-exitChannel:
				fmt.Println("Goroutinde 2 is done")
				return
			default:
				fmt.Println("loop 2")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	exitChannel <- true
	wg.Wait()

	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("Goroutinde 3 is done")
				return
			}
			fmt.Println("loop 3")
		}
	}()

	ch <- 1
	ch <- 1
	ch <- 1
	close(ch)
	wg.Wait()
}
