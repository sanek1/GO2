package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	workerCount := 1000
	count := 0

	for i := 0; i <= workerCount; i++ {
		wg.Add(1)
		go doit(i, done, &wg)
		count = i
	}

	close(done)
	wg.Wait()
	fmt.Printf("count = %v \n", count)
	fmt.Println("all done!")
}

func doit(workerId int, done <-chan struct{}, wg *sync.WaitGroup) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	for {
		select {
		case <-quit:
			return
		case <-done:
			fmt.Printf("[%v] is done\n", workerId)
			return
		}
	}

}
