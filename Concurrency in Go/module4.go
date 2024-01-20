// Implement the dining philosopher’s problem with the following constraints/modifications.

// 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

// 5. The host allows no more than 2 philosophers to eat concurrently.

// 6. Each philosopher is numbered, 1 through 5.

// 7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

// 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const numMeals = 3

var wg sync.WaitGroup
var hostSemaphore = make(chan struct{}, 2)
var chopsticks = make([]sync.Mutex, numPhilosophers)

type philosopher struct {
	id int
}

func (p philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(50))
}

func (p philosopher) eat() {
	fmt.Printf("Philosopher %d is starting to eat\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(100))
	fmt.Printf("Philosopher %d is finishing eating\n", p.id)
}

func (p philosopher) dine() {
	for i := 0; i < numMeals; i++ {
		p.think()

		// Ask for permission from the host
		hostSemaphore <- struct{}{}

		// Pick up left chopstick
		chopsticks[p.id].Lock()

		// Pick up right chopstick
		chopsticks[(p.id+1)%numPhilosophers].Lock()

		// Release permission from the host
		<-hostSemaphore

		p.eat()

		// Release chopsticks
		chopsticks[p.id].Unlock()
		chopsticks[(p.id+1)%numPhilosophers].Unlock()
	}
	wg.Done()
}

func main() {
	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosopher{id: i}.dine()
	}

	wg.Wait()
}
