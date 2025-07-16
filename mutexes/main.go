package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex

	rw sync.RWMutex 
)
func readCounter() int {
	rw.RLock() // Acquire a read lock
	defer rw.RUnlock() // Ensure the lock is released after reading
	return counter
}
func writeCounter() {
	rw.Lock() // Acquire a write lock
	defer rw.Unlock() // Ensure the lock is released after writing
	counter++
}
// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// mu.Lock() 
// 	counter++
// 	// mu.Unlock()

// }
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Printf("Reading Counter: %d\n", readCounter())
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			writeCounter()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}