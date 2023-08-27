package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type MutexMap struct {
	mx sync.RWMutex
	m  map[int]int
}

func NewCounters() *MutexMap {
	return &MutexMap{
		m: make(map[int]int),
	}
}

func (m *MutexMap) Read(key int) (int, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.m[key]
	return val, ok
}

func (m *MutexMap) Set(key int, value int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.m[key] = value
}

func main() {
	myMap := NewCounters()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap.Set(i, rand.Intn(100))
		}(i)
	}
	wg.Wait()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res, _ := myMap.Read(i)
			fmt.Println(res)
		}(i)
	}
	wg.Wait()

}
