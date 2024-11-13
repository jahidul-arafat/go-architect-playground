// race_test.go
package main

import (
	"fmt"
	"sync"
	"testing"
)

// Shared data structure for race condition testing
type Counter struct {
	Value int
}

// Function to increment the counter without any synchronization
func (c *Counter) IncrementUnsafe() {
	c.Value++
}

// Function to increment the counter with a mutex lock for synchronization
func (c *Counter) IncrementSafe(mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	c.Value++
}

func TestRaceConditionUnsafe(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Run 1000 goroutines that access the same counter without synchronization
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.IncrementUnsafe()
		}()
	}
	wg.Wait()

	fmt.Printf("Final counter value (unsafe): %d\n", counter.Value)
}

func TestRaceConditionSafe(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Run 1000 goroutines that access the same counter with synchronization
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.IncrementSafe(&mutex)
		}()
	}
	wg.Wait()

	fmt.Printf("Final counter value (safe): %d\n", counter.Value)
}
