package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var timeDuration int
	fmt.Println("Choose a timer")
	fmt.Fscan(os.Stdin, &timeDuration)

	channel := make(chan int, 1)
	timer := time.After(time.Second * time.Duration(timeDuration))

	for {
		select {
		case data := <-channel:
			fmt.Println(data)
		case channel <- rand.Intn(100):
		case <-timer:
			fmt.Println("Finish")
			return
		}
	}
}
