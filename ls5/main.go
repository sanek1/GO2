package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println(f1())
	fmt.Println(f2(1))
}

func f1() int64 {
	var c int64
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&c, 1)
		}()
	}
	time.Sleep(time.Second)
	return c
}

func f2(j int) int {
	var mu sync.Mutex
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			j += i
		}()
	}
	return j
}
