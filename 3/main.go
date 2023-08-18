package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mg sync.Mutex
	intArray := [5]int{2, 4, 6, 8, 10}
	var counter int

	for _, value := range intArray{
		wg.Add(1)
		go func(value int){
			defer wg.Done()
			mg.Lock()
			counter += int(math.Pow(float64(value), 2))
			mg.Unlock()
		}(value)
	}
	wg.Wait()
	fmt.Println(counter)
}