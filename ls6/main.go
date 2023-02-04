package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var counter int = 0

func main() {
	func1()
	func2()
}

func func2() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg := sync.WaitGroup{}
	for i := 0; i < 1<<4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {

			}
		}()
	}
	wg.Wait()
}

func func1() {
	ch := make(chan bool)
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The End")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock()
	ch <- true
}
