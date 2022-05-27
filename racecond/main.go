package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - sheesh.com")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.Lock()
		fmt.Println("one")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.Lock()
		fmt.Println("two")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.Lock()
		fmt.Println("three")
		score = append(score, 3)
		mut.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four ")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
