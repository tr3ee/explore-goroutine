package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100

func HeavyWork(id int) {
	rand.Seed(int64(id))
	interval := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(interval)
	fmt.Printf("HeavyWork %-3d cost %v\n", id, interval)
}

func main() {
	waitChan := make(chan int, 1)
	for i := 0; i < N; i++ {
		go func(n int) {
			HeavyWork(n)
			waitChan <- 1
		}(i)
	}
	cnt := 0
	for range waitChan {
		cnt++
		if cnt == N {
			break
		}
	}
	close(waitChan)
	fmt.Println("finished")
}
