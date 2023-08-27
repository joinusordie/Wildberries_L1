package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var countWorkers int
	fmt.Println("Select the number of workers")
	fmt.Fscan(os.Stdin, &countWorkers)

	workChannel := make(chan int)
	defer close(workChannel)
	exitSignal := make(chan os.Signal, 1)
	defer close(exitSignal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case data := <-workChannel:
					fmt.Fprintln(os.Stdout, data)
				case <-ctx.Done():
					fmt.Println("Worker finished")
					return
				}
			}
		}()
	}

	for {
		select {
		case workChannel <- rand.Intn(100):
		case <-exitSignal:
			cancel()
			wg.Wait()
			fmt.Println("Finish")
			return
		}
	}
}
