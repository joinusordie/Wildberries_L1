package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	intArray := [5]int{2, 4, 6, 8, 10}
	for _, value := range intArray {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, math.Pow(float64(i), 2))
		}(value)
	}
	wg.Wait()
}